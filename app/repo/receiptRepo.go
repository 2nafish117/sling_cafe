package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo_test/db"
	"sling_cafe/app/model"
	"sling_cafe/config"
	// "sling_cafe/util"
	"go.mongodb.org/mongo-driver/mongo"
	. "sling_cafe/log"
)

// ReceiptsFindAll /
// ReceiptsFindOne
// ReceiptsFindByEmpid
// ReceiptsFindOneById
// ReceiptsDeleteOne
// ReceiptsDeleteByEmpid
// ReceiptsUpdateOneById
// ReceiptsUpdateOne
// ReceiptsUpdateOneById
// ReceiptsUpdateOneByEmpid
// ReceiptsIsAlreadyExists
// ReceiptsIsAlreadyExistsWithId
// ReceiptsIsAlreadyExistsWithEmpid

// ReceiptsAggregate finds receipts based on pipeline and aggregation
func ReceiptsAggregate(ctx context.Context, pipeline interface{}) ([]*model.Receipt, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealsCollection)

	cursor, err := collection.Aggregate(ctx, pipeline)

	var receipts = make([]*model.Receipt, 0)
	if err != nil {
		Log.Error(err.Error())
		return receipts, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &receipts); err != nil {
		Log.Error(err.Error())
		return receipts, err
	}

	return receipts, nil
}

// ReceiptsFindAll finds all users receipts for all time
func ReceiptsFindAll(ctx context.Context) ([]*model.Receipt, error) {
	pipeline := mongo.Pipeline{
		// Stage 1, group by empid and add up the costs to find amtdue
		bson.D{primitive.E{
			Key: "$group", Value: bson.D{
				primitive.E{Key: "_id", Value: "$_id"}, // @TODO: what should _id of receipt be?
				primitive.E{Key: "empid", Value: bson.D{primitive.E{Key: "$first", Value: "$empid"}}},
				primitive.E{Key: "amtdue", Value: bson.D{primitive.E{Key: "$sum", Value: "$cost"}}},
			}},
		},
	}

	return ReceiptsAggregate(ctx, pipeline)
}
