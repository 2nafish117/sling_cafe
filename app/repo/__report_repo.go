package repo

// import (
// 	"context"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"sling_cafe/db"
// 	"sling_cafe/app/model"
// 	"sling_cafe/config"
// 	. "sling_cafe/log"
// 	"sling_cafe/util"
// )

// // ReportsAggregate /
// // ReportsFindOne /

// // ReportsAggregate finds receipts based on pipeline and aggregation
// func ReportsAggregate(ctx context.Context, pipeline mongo.Pipeline) ([]*model.Report, error) {
// 	conn := db.GetInstance()
// 	collection := conn.Database(config.GetInstance().DbName).Collection(MealsCollection)

// 	cursor, err := collection.Aggregate(ctx, pipeline)

// 	var reports = make([]*model.Report, 0)
// 	if err != nil {
// 		Log.Error(err.Error())
// 		return reports, err
// 	}
// 	defer cursor.Close(ctx)

// 	if err := cursor.All(ctx, &reports); err != nil {
// 		Log.Error(err.Error())
// 		return reports, err
// 	}

// 	return reports, nil
// }

// // @TODO: update func
// // BillsFindOne finds all users receipts for all time
// /*
// uses pileline
// db.meals.aggregate([
// 	{ $match: { datetime: { $gt: ISODate("2020-01-01T00:00:00.000Z"), $lte: ISODate("2020-07-01T00:00:00.000Z") } } },
// 	{ $project: { uid: 1, datetime: 1, type: 1, cost: 1, is_breakfast: {$cond: [{$eq: ['$type', "breakfast"]}, 1, 0]},cost_breakfast: {$cond: [{$eq: ['$type', "breakfast"]}, '$cost', 0]},is_lunch: {$cond: [{$eq: ['$type', "lunch"]}, 1, 0]},cost_lunch: {$cond: [{$eq: ['$type', "lunch"]}, '$cost', 0]},is_snack: {$cond: [{$eq: ['$type', "snack"]}, 1, 0]},cost_snack: {$cond: [{$eq: ['$type', "snack"]}, '$cost', 0]}}},
// 	{ $group: { _id: null, breakfast_quantity: { $sum: '$is_breakfast' }, breakfast_cost: { $sum: '$cost_breakfast' } ,lunch_quantity: { $sum: '$is_lunch' }, lunch_cost: { $sum: '$cost_lunch' }, snack_quantity: { $sum: '$is_snack' }, snack_cost: { $sum: '$cost_snack' }, total: { $sum: "$cost" }} },
// 	{ $project: { _id: 0, user: 1, total: 1, breakfast: { quantity: '$breakfast_quantity', cost: '$breakfast_cost' }, lunch: { quantity: '$lunch_quantity', cost: '$lunch_cost' }, snack: { quantity: '$snack_quantity', cost: '$snack_cost' }, }},
// ]).pretty()
// */
// func ReportsFindOne(ctx context.Context, fromDate, toDate primitive.DateTime) (*model.Report, error) {
// 	cp := util.CustomPipeline{}
// 	cp.Match(
// 		bson.D{
// 			{
// 				Key: "datetime",
// 				Value: bson.D{
// 					{Key: "$gt", Value: fromDate},
// 					{Key: "$lte", Value: toDate},
// 				},
// 			},
// 		},
// 	).Project(
// 		bson.D{
// 			{Key: "uid", Value: 1}, {Key: "datetime", Value: 1}, {Key: "type", Value: 1}, {Key: "cost", Value: 1},
// 			{Key: "is_breakfast", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$type", "breakfast"}}}, 1, 0}}}},
// 			{Key: "cost_breakfast", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$type", "breakfast"}}}, "$cost", 0}}}},
// 			{Key: "is_lunch", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$type", "lunch"}}}, 1, 0}}}},
// 			{Key: "cost_lunch", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$type", "lunch"}}}, "$cost", 0}}}},
// 			{Key: "is_snack", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$type", "snack"}}}, 1, 0}}}},
// 			{Key: "cost_snack", Value: bson.D{{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$type", "snack"}}}, "$cost", 0}}}},
// 		},
// 	).Group(
// 		bson.D{
// 			{Key: "_id", Value: nil},
// 			{Key: "breakfast_quantity", Value: bson.D{{Key: "$sum", Value: "$is_breakfast"}}},
// 			{Key: "breakfast_cost", Value: bson.D{{Key: "$sum", Value: "$cost_breakfast"}}},
// 			{Key: "lunch_quantity", Value: bson.D{{Key: "$sum", Value: "$is_lunch"}}},
// 			{Key: "lunch_cost", Value: bson.D{{Key: "$sum", Value: "$cost_lunch"}}},
// 			{Key: "snack_quantity", Value: bson.D{{Key: "$sum", Value: "$is_snack"}}},
// 			{Key: "snack_cost", Value: bson.D{{Key: "$sum", Value: "$cost_snack"}}},
// 			{Key: "total", Value: bson.D{{Key: "$sum", Value: "$cost"}}},
// 		},
// 	).Project(
// 		bson.D{
// 			{Key: "_id", Value: 0}, {Key: "total", Value: 1},
// 			{Key: "breakfast", Value: bson.D{
// 				{Key: "quantity", Value: "$breakfast_quantity"},
// 				{Key: "total", Value: "$breakfast_cost"},
// 			}},
// 			{Key: "lunch", Value: bson.D{
// 				{Key: "quantity", Value: "$lunch_quantity"},
// 				{Key: "total", Value: "$lunch_cost"},
// 			}},
// 			{Key: "snack", Value: bson.D{
// 				{Key: "quantity", Value: "$snack_quantity"},
// 				{Key: "total", Value: "$snack_cost"},
// 			}},
// 		},
// 	)

// 	report, err := ReportsAggregate(ctx, cp.Pipe)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(report) == 0 {
// 		return nil, nil
// 	}

// 	return report[0], nil
// }
