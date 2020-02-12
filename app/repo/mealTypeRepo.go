package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo_test/db"
	"sling_cafe/app/model"
	"sling_cafe/config"
)

// MealTypesInsertOne /
// MealTypesFindAll /
// MealTypesFindOne /
// MealTypesUpdateOne /

const mealTypesCollection string = "mealtypes"

// MealTypesInsertOne inserts one meal into the repository
func MealTypesInsertOne(ctx context.Context, mealType *model.MealType) (*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(mealTypesCollection)
	result, err := collection.InsertOne(ctx, mealType)

	if err != nil {
		return mealType, err
	}

	mealType.ID = result.InsertedID.(primitive.ObjectID)
	return mealType, nil
}

// MealTypesFindAll returns all meals
// @TODO pagination version of FindAll
func MealTypesFindAll(ctx context.Context) ([]*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(mealTypesCollection)
	cursor, err := collection.Find(ctx, bson.M{})
	mealTypes := make([]*model.MealType, 0)
	if err != nil {
		return mealTypes, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &mealTypes)
	return mealTypes, err
}

// MealTypesFindOne returns first meal matching query in repository
func MealTypesFindOne(ctx context.Context, query interface{}) (*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(mealTypesCollection)
	var mealType model.MealType
	err := collection.FindOne(ctx, query).Decode(&mealType)

	return &mealType, err
}

// MealTypesUpdateOne updates one meal in the repository and returns it
func MealTypesUpdateOne(ctx context.Context, mealType *model.MealType) (*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(mealTypesCollection)
	mt := &model.MealType{}
	err := collection.FindOneAndUpdate(ctx, bson.M{"_id": mealType.ID}, mealType).Decode(mt)
	return mt, err
}
