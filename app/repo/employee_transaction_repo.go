package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sling_cafe/app/model"
	"sling_cafe/config"
	"sling_cafe/db"
	// . "sling_cafe/log"
	"sling_cafe/util"
)

// EmployeeTransactionsCollection name of collection
const EmployeeTransactionsCollection string = "employee_transactions"

func EmployeeTransactionsFind(ctx context.Context, query interface{}) ([]*model.EmployeeTransaction, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(EmployeeTransactionsCollection)
	cursor, err := collection.Find(ctx, query)
	transactions := make([]*model.EmployeeTransaction, 0)
	if err != nil {
		return transactions, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &transactions); err != nil {
		return transactions, err
	}

	return transactions, nil
}

func EmployeeTransactionsFindAll(ctx context.Context) ([]*model.EmployeeTransaction, error) {
	return EmployeeTransactionsFind(ctx, bson.M{})
}

func EmployeeTransactionsFindOne(ctx context.Context, query interface{}) (*model.EmployeeTransaction, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(EmployeeTransactionsCollection)
	var transaction model.EmployeeTransaction
	err := collection.FindOne(ctx, query).Decode(&transaction)
	if err != nil {
		return &transaction, err
	}

	return &transaction, nil
}

func EmployeeTransactionsFindAllByEmployeeID(ctx context.Context, employeeID string) ([]*model.EmployeeTransaction, error) {
	return EmployeeTransactionsFind(ctx, bson.M{"employee_id": employeeID})
}

func EmployeeTransactionsInsertOne(ctx context.Context, transaction *model.EmployeeTransaction) (*model.EmployeeTransaction, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(EmployeeTransactionsCollection)
	result, err := collection.InsertOne(ctx, transaction)

	if err != nil {
		return transaction, err
	}
	transaction.ID = result.InsertedID.(primitive.ObjectID)
	return transaction, nil
}

func EmployeeTransactionsAggregate(ctx context.Context, pipeline mongo.Pipeline) (*mongo.Cursor, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(EmployeeTransactionsCollection)

	cursor, err := collection.Aggregate(ctx, pipeline)
	return cursor, err

	// var transactions = make([]*model.EmployeeTransaction, 0)
	// if err != nil {
	// 	Log.Error(err.Error())
	// 	return transactions, err
	// }
	// defer cursor.Close(ctx)

	// if err := cursor.All(ctx, &transactions); err != nil {
	// 	Log.Error(err.Error())
	// 	return transactions, err
	// }

	// return transactions, nil
}

// employee transaction report for all employees
// db.employee_transactions.aggregate([
// 	{ $project: { employee_id: 1, amount: {$cond: [{$eq: ['$transaction_type', "EmployeePayment"]}, {'$multiply':['$amount', NumberLong(-1)]}, '$amount']}}},
// 	{ $group: { _id: { employee_id: "$employee_id" }, amount: { $sum: '$amount' } } },
// 	{ $lookup: { from: "employees", localField: "_id.employee_id", foreignField: "employee_id", as: "employee" } },
// 	{ $unwind: { path: "$employee" } },
// 	{ $project: { _id: 0, name: '$employee.name', employee_id: '$employee.employee_id', amount: 1 }},
// ])

// EmployeeTransactionsAggregateReports
func EmployeeTransactionsAggregateReports(ctx context.Context) ([]*model.EmployeeTransactionReport, error) {
	cp := util.CustomPipeline{}
	cp.Project(
		bson.D{
			{Key: "employee_id", Value: 1},
			{Key: "amount", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$transaction_type", model.EmployeePayment}}}, bson.D{{Key: "$multiply", Value: bson.A{"$amount", int64(-1)}}}, "$amount"}}}},
		},
	).Group(
		bson.D{
			{Key: "_id", Value: bson.D{{Key: "employee_id", Value: "$employee_id"}}},
			{Key: "amount", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
		},
	).Lookup(
		bson.D{
			{Key: "from", Value: "employees"},
			{Key: "localField", Value: "_id.employee_id"},
			{Key: "foreignField", Value: "employee_id"},
			{Key: "as", Value: "employee"},
		},
	).Unwind(
		bson.D{
			{Key: "path", Value: "$employee"},
		},
	).Project(
		bson.D{
			{Key: "_id", Value: 0}, {Key: "employee_id", Value: "$employee.employee_id"}, {Key: "name", Value: "$employee.name"},
			{Key: "amount", Value: 1}, {Key: "workstation", Value: 1},
		},
	)

	var reports = make([]*model.EmployeeTransactionReport, 0)
	cursor, err := EmployeeTransactionsAggregate(context.TODO(), cp.Pipe)
	if err != nil {
		return reports, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &reports); err != nil {
		return reports, err
	}

	return reports, nil
}

// db.employee_transactions.aggregate([
// 	{ $match: { employee_id: "4" }},
// 	{ $project: { employee_id: 1, amount: {$cond: [{$eq: ['$transaction_type', "EmployeePayment"]}, {'$multiply':['$amount', NumberLong(-1)]}, '$amount']}}},
// 	{ $group: { _id: { employee_id: "$employee_id" }, amount: { $sum: '$amount' } } },
// 	{ $lookup: { from: "employees", localField: "_id.employee_id", foreignField: "employee_id", as: "employee" } },
// 	{ $unwind: { path: "$employee" } },
// 	{ $project: { _id: 0, name: '$employee.name', employee_id: '$employee.employee_id', amount: 1 }},
// ])

// EmployeeTransactionsAggregateReportByEmployeeID
func EmployeeTransactionsAggregateReportByEmployeeID(ctx context.Context, employeeID string) (*model.EmployeeTransactionReport, error) {
	cp := util.CustomPipeline{}
	cp.Match(
		bson.D{{Key: "employee_id", Value: employeeID}},
	).Project(
		bson.D{
			{Key: "employee_id", Value: 1},
			{Key: "amount", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$transaction_type", model.EmployeePayment}}}, bson.D{{Key: "$multiply", Value: bson.A{"$amount", int64(-1)}}}, "$amount"}}}},
		},
	).Group(
		bson.D{
			{Key: "_id", Value: bson.D{{Key: "employee_id", Value: "$employee_id"}}},
			{Key: "amount", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
		},
	).Lookup(
		bson.D{
			{Key: "from", Value: "employees"},
			{Key: "localField", Value: "_id.employee_id"},
			{Key: "foreignField", Value: "employee_id"},
			{Key: "as", Value: "employee"},
		},
	).Unwind(
		bson.D{
			{Key: "path", Value: "$employee"},
		},
	).Project(
		bson.D{
			{Key: "_id", Value: 0}, {Key: "employee_id", Value: "$employee.employee_id"}, {Key: "name", Value: "$employee.name"},
			{Key: "amount", Value: 1},
		},
	)

	var reports = make([]*model.EmployeeTransactionReport, 0)
	cursor, err := EmployeeTransactionsAggregate(context.TODO(), cp.Pipe)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &reports); err != nil {
		return nil, err
	}

	if len(reports) == 0 {
		return nil, nil
	}

	return reports[0], nil
}
