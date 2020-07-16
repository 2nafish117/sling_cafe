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
	"time"
	// . "sling_cafe/log"
	// "fmt"
)

// MealEntriesCollection name of collection
const MealEntriesCollection string = "meal_entries"

func MealEntriesFind(ctx context.Context, query interface{}) ([]*model.MealEntry, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealEntriesCollection)
	cursor, err := collection.Find(ctx, query)
	meals := make([]*model.MealEntry, 0)
	if err != nil {
		return meals, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &meals); err != nil {
		return meals, err
	}

	return meals, nil
}

func MealEntriesFindAll(ctx context.Context) ([]*model.MealEntry, error) {
	return MealEntriesFind(ctx, bson.M{})
}

func MealEntriesFindOne(ctx context.Context, query interface{}) (*model.MealEntry, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealEntriesCollection)
	var meal model.MealEntry
	err := collection.FindOne(ctx, query).Decode(&meal)
	if err != nil {
		return &meal, err
	}

	return &meal, nil
}

func MealEntriesFindAllByEmployeeID(ctx context.Context, employeeID string) ([]*model.MealEntry, error) {
	return MealEntriesFind(ctx, bson.M{"employee_id": employeeID})
}

func MealEntriesHasEmployeeAlreadyEaten(ctx context.Context, t time.Time, employeeID, mealID model.MealID) bool {
	meals, err := MealEntriesFind(ctx, bson.M{"employee_id": employeeID, "meal_id": mealID})

	isOnSameDay := false
	for _, meal := range meals {
		year, month, day := meal.DateTime.Date()
		y, m, d := t.Date()
		isOnSameDay = isOnSameDay || (year == y && month == m && day == d)
	}
	// if a meal with employee_id and meal_id has been found AND the dates match,
	// then that person has eaten that meal on that day
	// fmt.Println(err == nil)
	// fmt.Println(isOnSameDay)
	return err == nil && isOnSameDay
}

// MealEntriesInsertOne inserts one meal into the repository
func MealEntriesInsertOne(ctx context.Context, meal *model.MealEntry) (*model.MealEntry, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealEntriesCollection)
	result, err := collection.InsertOne(ctx, meal)

	if err != nil {
		return meal, err
	}
	meal.ID = result.InsertedID.(primitive.ObjectID)
	return meal, nil
}

func MealEntriesAggregate(ctx context.Context, pipeline mongo.Pipeline) (*mongo.Cursor, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealEntriesCollection)

	cursor, err := collection.Aggregate(ctx, pipeline)
	return cursor, err

	// var meals = make([]*model.MealEntryReport, 0)
	// if err != nil {
	// 	Log.Error(err.Error())
	// 	return meals, err
	// }
	// defer cursor.Close(ctx)

	// if err := cursor.All(ctx, &meals); err != nil {
	// 	Log.Error(err.Error())
	// 	return meals, err
	// }

	// return meals, nil
}

// db.meal_entries.aggregate([
// 	{ $project: { employee_id: 1, is_breakfast: {$cond: [{$eq: ['$meal_id', "breakfast"]}, 1, 0]}, is_lunch: {$cond: [{$eq: ['$meal_id', "lunch"]}, 1, 0]}, is_snack: {$cond: [{$eq: ['$meal_id', "snack"]}, 1, 0]}}},
// 	{ $group: { _id: { employee_id: "$employee_id" }, breakfast_quantity: { $sum: '$is_breakfast' }, lunch_quantity: { $sum: '$is_lunch' }, snack_quantity: { $sum: '$is_snack' } } },
// 	{ $lookup: { from: "employees", localField: "_id.employee_id", foreignField: "employee_id", as: "employee" } },
// 	{ $unwind: { path: "$employee" } },
// 	{ $project: { _id: 0, name: '$employee.name', employee_id: '$employee.employee_id', breakfast_quantity: '$breakfast_quantity', lunch_quantity: '$lunch_quantity', snack_quantity: '$snack_quantity' }},
// ]).pretty()
// MealEntriesAggregateReports
func MealEntriesAggregateReports(ctx context.Context) ([]*model.MealEntryReport, error) {
	cp := util.CustomPipeline{}
	cp.Project(
		bson.D{
			{Key: "employee_id", Value: 1},
			{Key: "is_breakfast", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$meal_id", model.Breakfast}}}, 1, 0}}}},
			{Key: "is_lunch", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$meal_id", model.Lunch}}}, 1, 0}}}},
			{Key: "is_snack", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$meal_id", model.Snack}}}, 1, 0}}}},
		},
	).Group(
		bson.D{
			{Key: "_id", Value: bson.D{{Key: "employee_id", Value: "$employee_id"}}},
			{Key: "breakfast_quantity", Value: bson.D{{Key: "$sum", Value: "$is_breakfast"}}},
			{Key: "lunch_quantity", Value: bson.D{{Key: "$sum", Value: "$is_lunch"}}},
			{Key: "snack_quantity", Value: bson.D{{Key: "$sum", Value: "$is_snack"}}},
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
			{Key: "breakfast_quantity", Value: "$breakfast_quantity"},
			{Key: "lunch_quantity", Value: "$lunch_quantity"},
			{Key: "snack_quantity", Value: "$snack_quantity"},
		},
	)

	var reports = make([]*model.MealEntryReport, 0)
	cursor, err := MealEntriesAggregate(context.TODO(), cp.Pipe)
	if err != nil {
		return reports, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &reports); err != nil {
		return reports, err
	}

	return reports, nil
}

// db.meal_entries.aggregate([
// 	{ $match: { employee_id: empid }},
// 	{ $project: { employee_id: 1, is_breakfast: {$cond: [{$eq: ['$meal_id', "breakfast"]}, 1, 0]}, is_lunch: {$cond: [{$eq: ['$meal_id', "lunch"]}, 1, 0]}, is_snack: {$cond: [{$eq: ['$meal_id', "snack"]}, 1, 0]}}},
// 	{ $group: { _id: { employee_id: "$employee_id" }, breakfast_quantity: { $sum: '$is_breakfast' }, lunch_quantity: { $sum: '$is_lunch' }, snack_quantity: { $sum: '$is_snack' } } },
// 	{ $lookup: { from: "employees", localField: "_id.employee_id", foreignField: "employee_id", as: "employee" } },
// 	{ $unwind: { path: "$employee" } },
// 	{ $sort: { total: -1 } },
// 	{ $project: { _id: 0, name: '$employee.name', employee_id: '$employee.employee_id', breakfast_quantity: '$breakfast_quantity', lunch_quantity: '$lunch_quantity', snack_quantity: '$snack_quantity' }},
// ]).pretty()
// MealEntriesAggregateReportByEmployeeID
func MealEntriesAggregateReportByEmployeeID(ctx context.Context, employeeID string) (*model.MealEntryReport, error) {
	cp := util.CustomPipeline{}
	cp.Match(
		bson.D{{Key: "employee_id", Value: employeeID}},
	).Project(
		bson.D{
			{Key: "employee_id", Value: 1},
			{Key: "is_breakfast", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$meal_id", model.Breakfast}}}, 1, 0}}}},
			{Key: "is_lunch", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$meal_id", model.Lunch}}}, 1, 0}}}},
			{Key: "is_snack", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$meal_id", model.Snack}}}, 1, 0}}}},
		},
	).Group(
		bson.D{
			{Key: "_id", Value: bson.D{{Key: "employee_id", Value: "$employee_id"}}},
			{Key: "breakfast_quantity", Value: bson.D{{Key: "$sum", Value: "$is_breakfast"}}},
			{Key: "lunch_quantity", Value: bson.D{{Key: "$sum", Value: "$is_lunch"}}},
			{Key: "snack_quantity", Value: bson.D{{Key: "$sum", Value: "$is_snack"}}},
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
			{Key: "breakfast_quantity", Value: "$breakfast_quantity"},
			{Key: "lunch_quantity", Value: "$lunch_quantity"},
			{Key: "snack_quantity", Value: "$snack_quantity"},
		},
	)

	var reports = make([]*model.MealEntryReport, 0)
	cursor, err := MealEntriesAggregate(context.TODO(), cp.Pipe)
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
