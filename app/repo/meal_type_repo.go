package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"mongo_test/db"
	"sling_cafe/app/model"
	"sling_cafe/config"
	"sling_cafe/util"
)

// MealTypesFindAll /
// MealTypesFindOne /
// MealTypesFindOneById /
// MealTypesDeleteOne /
// MealTypesUpdateOneById /
// MealTypesUpdateOne /
// MealTypesUpdateOneById /
// MealTypesIsAlreadyExists /
// MealTypesIsAlreadyExistsWithId /

// MealTypesCollection name of users collection
const MealTypesCollection string = "mealtypes"

// MealTypesFindAll returns all mealtypes from repo
// @TODO pagination version of FindAll
func MealTypesFindAll(ctx context.Context) ([]*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealTypesCollection)

	cursor, err := collection.Find(ctx, bson.M{})

	var mealtypes []*model.MealType
	if err != nil {
		return mealtypes, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &mealtypes); err != nil {
		return mealtypes, err
	}

	return mealtypes, nil
}

// MealtypesFindOne finds first mealtype matching query and returns it
func MealtypesFindOne(ctx context.Context, query interface{}) (*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealTypesCollection)
	var mealtype model.MealType
	err := collection.FindOne(ctx, query).Decode(&mealtype)
	return &mealtype, err
}

// MealTypesFindOneByType find the mealtype by id
func MealTypesFindOneByType(ctx context.Context, typ string) (*model.MealType, error) {
	return MealtypesFindOne(ctx, bson.M{"type": typ})
}

// MealTypesInsertOne inserts one mealtype to the repository
func MealTypesInsertOne(ctx context.Context, mt *model.MealType) (*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealTypesCollection)
	result, err := collection.InsertOne(ctx, mt)

	if err != nil {
		return mt, err
	}
	mt.ID = result.InsertedID.(primitive.ObjectID)
	return mt, nil
}

// MealTypesDeleteOne deletes a mealtype and returns it
func MealTypesDeleteOne(ctx context.Context, query interface{}) (*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealTypesCollection)
	var mt model.MealType
	err := collection.FindOneAndDelete(ctx, query).Decode(&mt)
	return &mt, err
}

func MealTypesDeleteOneByType(ctx context.Context, typ string) (*model.MealType, error) {
	return MealTypesDeleteOne(ctx, bson.M{"type": typ})
}

// MealTypesUpdateOne updates one
func MealTypesUpdateOne(ctx context.Context, query interface{}, update interface{}) (*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealTypesCollection)
	var mt model.MealType
	err := collection.FindOneAndUpdate(ctx, query, update).Decode(&mt)
	if err != nil {
		return nil, err
	}
	return &mt, nil
}

// MealTypesUpdateOneByType updates one by id
func MealTypesUpdateOneByType(ctx context.Context, typ string, update *model.MealType) (*model.MealType, error) {
	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}

	return MealTypesUpdateOne(ctx, bson.M{"type": typ}, upd)
}

// MealTypesIsAlreadyExists asks repo if melatype already exists
func MealTypesIsAlreadyExists(ctx context.Context, query interface{}) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
	var mt model.MealType
	erro := collection.FindOne(ctx, query).Decode(&mt)

	return erro == nil
}

// MealTypesIsAlreadyExistsWithId asks repo if already exists, by id
func MealTypesIsAlreadyExistsWithType(ctx context.Context, typ string) bool {
	return MealTypesIsAlreadyExists(ctx, bson.M{"type": typ})
}
