package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/app/model"
	"sling_cafe/config"
	"sling_cafe/db"
	"sling_cafe/util"
)

// CaterersCollection name of vendors collection
const CaterersCollection string = "caterers"

func CaterersFind(ctx context.Context, query interface{}) ([]*model.Caterer, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(CaterersCollection)

	cursor, err := collection.Find(ctx, query)

	caterers := make([]*model.Caterer, 0)
	if err != nil {
		return caterers, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &caterers); err != nil {
		return caterers, err
	}

	return caterers, nil
}

func CaterersFindAll(ctx context.Context) ([]*model.Caterer, error) {
	return CaterersFind(ctx, bson.M{})
}

func CaterersFindOne(ctx context.Context, query interface{}) (*model.Caterer, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(CaterersCollection)
	var caterer model.Caterer
	err := collection.FindOne(ctx, query).Decode(&caterer)
	return &caterer, err
}

func CaterersFindByCatererID(ctx context.Context, catererID string) (*model.Caterer, error) {
	return CaterersFindOne(ctx, bson.M{"caterer_id": catererID})
}

func CaterersFindOneByID(ctx context.Context, id string) (*model.Caterer, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return CaterersFindOne(ctx, bson.M{"_id": internalID})
}

func CaterersInsertOne(ctx context.Context, c *model.Caterer) (*model.Caterer, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(CaterersCollection)
	result, err := collection.InsertOne(ctx, c)

	if err != nil {
		return c, err
	}
	c.ID = result.InsertedID.(primitive.ObjectID)
	return c, nil
}

func CaterersDeleteOne(ctx context.Context, query interface{}) (*model.Caterer, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(CaterersCollection)
	var c model.Caterer
	err := collection.FindOneAndDelete(ctx, query).Decode(&c)
	return &c, err
}

func CaterersDeleteByCatererID(ctx context.Context, catererID string) (*model.Caterer, error) {
	return CaterersDeleteOne(ctx, bson.M{"caterer_id": catererID})
}

func CaterersDeleteByID(ctx context.Context, id string) (*model.Caterer, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return CaterersDeleteOne(ctx, bson.M{"_id": internalID})
}

func CaterersUpdateOne(ctx context.Context, query interface{}, update interface{}) (*model.Caterer, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(CaterersCollection)
	var c model.Caterer
	err := collection.FindOneAndUpdate(ctx, query, update).Decode(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func CaterersUpdateOneByID(ctx context.Context, id string, update *model.Caterer) (*model.Caterer, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}

	return CaterersUpdateOne(ctx, bson.M{"_id": internalID}, upd)
}

func CaterersUpdateOneByCatererID(ctx context.Context, CatererID string, update *model.Caterer) (*model.Caterer, error) {
	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}
	return CaterersUpdateOne(ctx, bson.M{"caterer_id": CatererID}, upd)
}

func CaterersIsAlreadyExists(ctx context.Context, query interface{}) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(CaterersCollection)
	var c model.Caterer
	erro := collection.FindOne(ctx, query).Decode(&c)

	return erro == nil
}

func CaterersIsAlreadyExistsWithID(ctx context.Context, id string) bool {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	return CaterersIsAlreadyExists(ctx, bson.M{"_id": internalID})
}

func CaterersIsAlreadyExistsWithCatererID(ctx context.Context, catererID string) bool {
	return CaterersIsAlreadyExists(ctx, bson.M{"caterer_id": catererID})
}

func CaterersIsInactive(ctx context.Context, catererID string) bool {
	cat, err := CaterersFindByCatererID(ctx, catererID)
	if err != nil {
		return false
	}
	return cat.Inactive
}
