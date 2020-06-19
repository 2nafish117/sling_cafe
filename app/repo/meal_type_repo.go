package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"errors"
	"mongo_test/db"
	"sling_cafe/app/model"
	"sling_cafe/config"
	"sling_cafe/util"
	"time"
)

// MealTypesCollection name of users collection
const MealTypesCollection string = "meal_types"

func MealTypesFind(ctx context.Context, query interface{}) ([]*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealTypesCollection)

	cursor, err := collection.Find(ctx, query)

	mealtypes := make([]*model.MealType, 0)
	if err != nil {
		return mealtypes, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &mealtypes); err != nil {
		return mealtypes, err
	}

	return mealtypes, nil
}

func MealTypesFindAll(ctx context.Context) ([]*model.MealType, error) {
	return MealTypesFind(ctx, bson.M{})
}

// MealtypesFindOne finds first mealtype matching query and returns it
func MealTypesFindOne(ctx context.Context, query interface{}) (*model.MealType, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealTypesCollection)
	var mealtype model.MealType
	err := collection.FindOne(ctx, query).Decode(&mealtype)
	return &mealtype, err
}

func MealTypesFindOneByTimeOfDay(ctx context.Context, dt time.Time) (*model.MealType, error) {
	mealTypes, err := MealTypesFindAll(ctx)
	if err != nil {
		return nil, err
	}
	hour, minute, second := dt.Clock()
	convertedTime := time.Date(2020, 1, 1, hour, minute, second, 0, dt.Location())

	for _, mt := range mealTypes {
		if convertedTime.Unix() > mt.FromTime.Unix() && convertedTime.Unix() < mt.ToTime.Unix() {
			return mt, nil
		}
	}

	return nil, errors.New("no active mealtype at this time, come back later")
}

// MealTypesFindOneByType find the mealtype by id
func MealTypesFindOneByMealID(ctx context.Context, mealID string) (*model.MealType, error) {
	return MealTypesFindOne(ctx, bson.M{"meal_id": mealID})
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

func MealTypesDeleteOneByMealID(ctx context.Context, mealID string) (*model.MealType, error) {
	return MealTypesDeleteOne(ctx, bson.M{"meal_id": mealID})
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
func MealTypesUpdateOneByMealID(ctx context.Context, mealID string, update *model.MealType) (*model.MealType, error) {
	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}

	return MealTypesUpdateOne(ctx, bson.M{"meal_id": mealID}, upd)
}

// MealTypesIsAlreadyExists asks repo if melatype already exists
func MealTypesIsAlreadyExists(ctx context.Context, query interface{}) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(MealTypesCollection)
	var mt model.MealType
	erro := collection.FindOne(ctx, query).Decode(&mt)

	return erro == nil
}

// MealTypesIsAlreadyExistsWithId asks repo if already exists, by id
func MealTypesIsAlreadyExistsWithMealID(ctx context.Context, mealID model.MealID) bool {
	return MealTypesIsAlreadyExists(ctx, bson.M{"meal_id": mealID})
}
