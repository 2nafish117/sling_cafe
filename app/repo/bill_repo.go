package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mongo_test/db"
	"sling_cafe/app/model"
	"sling_cafe/config"
	. "sling_cafe/log"
	"sling_cafe/util"
)

// BillsAggregate /
// BillsFindAll
// BillsFindOne

// BillsAggregate finds receipts based on pipeline and aggregation
func BillsAggregate(ctx context.Context, pipeline mongo.Pipeline) ([]*model.Bill, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealsCollection)

	cursor, err := collection.Aggregate(ctx, pipeline)

	var bills = make([]*model.Bill, 0)
	if err != nil {
		Log.Error(err.Error())
		return bills, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &bills); err != nil {
		Log.Error(err.Error())
		return bills, err
	}

	return bills, nil
}

// BillsFindAll finds all users receipts for all time
// sort -1 for descending, 1 for ascending
func BillsFindAll(ctx context.Context, fromDate, toDate primitive.DateTime, sort int) ([]*model.Bill, error) {
	cp := util.CustomPipeline{}
	cp.Match(
		bson.D{
			{
				Key: "datetime",
				Value: bson.D{
					{Key: "$gt", Value: fromDate},
					{Key: "$lte", Value: toDate},
				},
			},
		},
	).Group(
		bson.D{
			{Key: "_id", Value: bson.D{{Key: "uid", Value: "$uid"}}}, {Key: "total", Value: bson.D{{Key: "$sum", Value: "$cost"}}},
		},
	).Lookup(
		bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "_id.uid"},
			{Key: "foreignField", Value: "uid"},
			{Key: "as", Value: "user"},
		},
	).Unwind(
		bson.D{
			{Key: "path", Value: "$user"},
		},
	).Project(
		bson.D{
			{Key: "_id", Value: 0},
		},
	).Sort(
		bson.D{
			{Key: "total", Value: sort},
		},
	)

	return BillsAggregate(context.TODO(), cp.Pipe)
}

// @TODO: update func
// BillsFindOne finds all users receipts for all time
func BillsFindOne(ctx context.Context, uid string, fromDate, toDate primitive.DateTime) (*model.Bill, error) {
	cp := util.CustomPipeline{}
	cp.Match(
		bson.D{
			{Key: "uid", Value: uid},
			{Key: "datetime",
				Value: bson.D{
					{Key: "$gt", Value: fromDate},
					{Key: "$lte", Value: toDate},
				}},
		},
	).Group(
		bson.D{
			{Key: "_id",
				Value: bson.D{
					{Key: "uid", Value: "$uid"},
				}},
			{Key: "total",
				Value: bson.D{
					{Key: "$sum", Value: "$cost"},
				}},
		},
	).Lookup(
		bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "_id.uid"},
			{Key: "foreignField", Value: "uid"},
			{Key: "as", Value: "user"},
		},
	).Unwind(
		bson.D{
			{Key: "path", Value: "$user"},
		},
	).Project(
		bson.D{
			{Key: "_id", Value: 0},
		},
	)

	bill, err := BillsAggregate(ctx, cp.Pipe)
	if err != nil {
		return nil, err
	}
	return bill[0], nil
}
