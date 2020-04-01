package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo_test/db"
	"sling_cafe/app/model"
	"sling_cafe/config"
)

// PaymentsFindAll /
// PaymentsFindOne /
// PaymentsFindAllByUId /
// PaymentsInsertOne /

// MealsCollection name of collection
const PaymentsCollection string = "payments"

// PaymentsFindAll returns all payments
// @TODO pagination version of FindAll
func PaymentsFindAll(ctx context.Context) ([]*model.Payment, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(PaymentsCollection)
	cursor, err := collection.Find(ctx, bson.M{})
	payments := make([]*model.Payment, 0)
	if err != nil {
		return payments, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &payments); err != nil {
		return payments, err
	}

	return payments, nil
}

// PaymentsFindOne returns first meal matching query
func PaymentsFindOne(ctx context.Context, query interface{}) (*model.Payment, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(PaymentsCollection)
	var payment model.Payment
	err := collection.FindOne(ctx, query).Decode(&payment)
	if err != nil {
		return &payment, err
	}

	return &payment, nil
}

// PaymentsFindAllByUId finds all the payments by an employee of uid
func PaymentsFindAllByUId(ctx context.Context, uid string) ([]*model.Payment, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(PaymentsCollection)

	cursor, err := collection.Find(ctx, bson.M{"uid": uid})
	payments := make([]*model.Payment, 0)
	if err != nil {
		return payments, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &payments); err != nil {
		return payments, err
	}

	return payments, nil
}

// PaymentsInsertOne inserts one meal into the repository
func PaymentsInsertOne(ctx context.Context, payment *model.Payment) (*model.Payment, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(PaymentsCollection)
	result, err := collection.InsertOne(ctx, payment)

	if err != nil {
		return payment, err
	}
	payment.ID = result.InsertedID.(primitive.ObjectID)
	return payment, nil
}
