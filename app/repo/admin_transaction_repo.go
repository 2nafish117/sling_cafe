package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sling_cafe/app/model"
	"sling_cafe/config"
	"sling_cafe/db"
	"sling_cafe/util"
)

// AdminTransactionsCollection name of collection
const AdminTransactionsCollection string = "admin_transactions"

func AdminTransactionsFind(ctx context.Context, query interface{}) ([]*model.AdminTransaction, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(AdminTransactionsCollection)
	cursor, err := collection.Find(ctx, query)
	transactions := make([]*model.AdminTransaction, 0)
	if err != nil {
		return transactions, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &transactions); err != nil {
		return transactions, err
	}

	return transactions, nil
}

func AdminTransactionsFindAll(ctx context.Context) ([]*model.AdminTransaction, error) {
	return AdminTransactionsFind(ctx, bson.M{})
}

func AdminTransactionsFindOne(ctx context.Context, query interface{}) (*model.AdminTransaction, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(AdminTransactionsCollection)
	var transaction model.AdminTransaction
	err := collection.FindOne(ctx, query).Decode(&transaction)
	if err != nil {
		return &transaction, err
	}

	return &transaction, nil
}

func AdminTransactionsFindAllByAdminID(ctx context.Context, adminID string) ([]*model.AdminTransaction, error) {
	return AdminTransactionsFind(ctx, bson.M{"admin_id": adminID})
}

func AdminTransactionsInsertOne(ctx context.Context, transaction *model.AdminTransaction) (*model.AdminTransaction, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(AdminTransactionsCollection)
	result, err := collection.InsertOne(ctx, transaction)

	if err != nil {
		return transaction, err
	}
	transaction.ID = result.InsertedID.(primitive.ObjectID)
	return transaction, nil
}

func AdminTransactionsAggregate(ctx context.Context, pipeline mongo.Pipeline) (*mongo.Cursor, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(AdminTransactionsCollection)

	cursor, err := collection.Aggregate(ctx, pipeline)
	return cursor, err

	// var transactions = make([]*model.AdminTransaction, 0)
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

// admin transaction report for all admins
// db.employee_transactions.aggregate([
// 	{ $project: { admin_id: 1, amount: {$cond: [{$eq: ['$transaction_type', "DepositAmount"]}, {'$multiply':['$amount', NumberLong(-1)]}, '$amount']}}},
// 	{ $group: { _id: { admin_id: "$admin_id" }, amount: { $sum: '$amount' } } },
// 	{ $lookup: { from: "admins", localField: "_id.admin_id", foreignField: "admin_id", as: "admin" } },
// 	{ $unwind: { path: "$admin" } },
// 	{ $project: { _id: 0, name: '$admin.name', admin_id: '$admin.admin_id', amount: 1 }},
// ])

// AdminTransactionsAggregateReports
func AdminTransactionsAggregateReports(ctx context.Context) ([]*model.AdminTransactionReport, error) {
	cp := util.CustomPipeline{}
	cp.Project(
		bson.D{
			{Key: "admin_id", Value: 1},
			{Key: "amount", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$transaction_type", model.AdminDeposit}}}, bson.D{{Key: "$multiply", Value: bson.A{"$amount", int64(-1)}}}, "$amount"}}}},
		},
	).Group(
		bson.D{
			{Key: "_id", Value: bson.D{{Key: "admin_id", Value: "$admin_id"}}},
			{Key: "amount", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
		},
	).Lookup(
		bson.D{
			{Key: "from", Value: "admins"},
			{Key: "localField", Value: "_id.admin_id"},
			{Key: "foreignField", Value: "admin_id"},
			{Key: "as", Value: "admin"},
		},
	).Unwind(
		bson.D{
			{Key: "path", Value: "$admin"},
		},
	).Project(
		bson.D{
			{Key: "_id", Value: 0}, {Key: "admin_id", Value: "$admin.admin_id"}, {Key: "name", Value: "$admin.name"},
			{Key: "amount", Value: 1},
		},
	)

	var reports = make([]*model.AdminTransactionReport, 0)
	cursor, err := AdminTransactionsAggregate(context.TODO(), cp.Pipe)
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
// 	{ $match: { admin_id: "4" }},
// 	{ $project: { admin_id: 1, amount: {$cond: [{$eq: ['$transaction_type', "EmployeePayment"]}, {'$multiply':['$amount', NumberLong(-1)]}, '$amount']}}},
// 	{ $group: { _id: { admin_id: "$admin_id" }, amount: { $sum: '$amount' } } },
// 	{ $lookup: { from: "admins", localField: "_id.admin_id", foreignField: "admin_id", as: "admin" } },
// 	{ $unwind: { path: "$admin" } },
// 	{ $project: { _id: 0, name: '$admin.name', admin_id: '$admin.admin_id', amount: 1 }},
// ])

// AdminTransactionsAggregateReportByAdminID
func AdminTransactionsAggregateReportByAdminID(ctx context.Context, employeeID string) (*model.AdminTransactionReport, error) {
	cp := util.CustomPipeline{}
	cp.Match(
		bson.D{{Key: "admin_id", Value: employeeID}},
	).Project(
		bson.D{
			{Key: "admin_id", Value: 1},
			{Key: "amount", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$transaction_type", model.EmployeePayment}}}, bson.D{{Key: "$multiply", Value: bson.A{"$amount", int64(-1)}}}, "$amount"}}}},
		},
	).Group(
		bson.D{
			{Key: "_id", Value: bson.D{{Key: "admin_id", Value: "$admin_id"}}},
			{Key: "amount", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
		},
	).Lookup(
		bson.D{
			{Key: "from", Value: "admins"},
			{Key: "localField", Value: "_id.admin_id"},
			{Key: "foreignField", Value: "admin_id"},
			{Key: "as", Value: "admin"},
		},
	).Unwind(
		bson.D{
			{Key: "path", Value: "$admin"},
		},
	).Project(
		bson.D{
			{Key: "_id", Value: 0}, {Key: "admin_id", Value: "$admin.admin_id"}, {Key: "name", Value: "$admin.name"},
			{Key: "amount", Value: 1},
		},
	)

	var reports = make([]*model.AdminTransactionReport, 0)
	cursor, err := AdminTransactionsAggregate(context.TODO(), cp.Pipe)
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
