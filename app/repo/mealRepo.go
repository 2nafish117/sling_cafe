package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo_test/db"
	"sling_cafe/app/model"
	"sling_cafe/config"
)

// MealsFindAll /
// MealsFindOne /
// MealsFindAllByEmpid /
// MealsInsertOne /

const mealsCollection string = "meals"

// @TODO pagination version of FindAll
func MealsFindAll(ctx context.Context) ([]*model.Meal, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(mealsCollection)
	cursor, err := collection.Find(ctx, bson.M{})
	meals := make([]*model.Meal, 0)
	if err != nil {
		return meals, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &meals); err != nil {
		return meals, err
	}

	return meals, nil
}

func MealsFindOne(ctx context.Context, query interface{}) (*model.Meal, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(mealsCollection)
	var meal model.Meal
	err := collection.FindOne(ctx, query).Decode(&meal)
	if err != nil {
		return &meal, err
	}

	return &meal, nil
}

func MealsFindAllByEmpid(ctx context.Context, empid string) ([]*model.Meal, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(mealsCollection)
	cursor, err := collection.Find(ctx, bson.M{"empid": empid})
	meals := make([]*model.Meal, 0)
	if err != nil {
		return meals, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &meals); err != nil {
		return meals, err
	}

	return meals, nil
}

func MealsInsertOne(ctx context.Context, meal *model.Meal) (*model.Meal, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(mealsCollection)
	result, err := collection.InsertOne(ctx, meal)

	if err != nil {
		return meal, err
	}
	meal.ID = result.InsertedID.(primitive.ObjectID)
	return meal, nil
}
