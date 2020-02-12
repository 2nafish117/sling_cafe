package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo_test/db"
	"sling_cafe/app/model"
	"sling_cafe/config"
)

// UsersFindAll /
// UsersFindOne /
// UsersFindByEmpid /
// UsersInsertOne /
// UsersDeleteOne /
// UsersDeleteByEmpid /

const usersCollection string = "users"

// @TODO pagination version of FindAll
func UsersFindAll(ctx context.Context) ([]*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
	cursor, err := collection.Find(ctx, bson.M{})
	var users []*model.User
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &users); err != nil {
		return users, err
	}

	return users, nil
}

func UsersFindOne(ctx context.Context, query interface{}) (*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
	var user model.User
	err := collection.FindOne(ctx, query).Decode(&user)
	return &user, err
}

func UsersFindByEmpid(ctx context.Context, empid string) (*model.User, error) {
	return UsersFindOne(ctx, bson.M{"empid": empid})
}

// FindOneById, find the user by the provided id
// return matched user and error if any
func UsersFindOneById(ctx context.Context, id string) (*model.User, error) {
	internalId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return UsersFindOne(ctx, bson.M{"_id": internalId})
}

// Create, will perform db opration to save user
// Returns modified user and error if occurs
func UsersInsertOne(ctx context.Context, u *model.User) (*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
	result, err := collection.InsertOne(ctx, u)

	if err != nil {
		return u, err
	}
	u.ID = result.InsertedID.(primitive.ObjectID)
	return u, nil
}

// Delete, will remove user entry from DB
// Return error if any
func UsersDeleteOne(ctx context.Context, user *model.User) (*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
	var u model.User
	err := collection.FindOneAndDelete(ctx, bson.M{"_id": user.ID}).Decode(&u)
	fmt.Print(u, err)
	return &u, err
}

func UsersDeleteByEmpid(ctx context.Context, empid string) (*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
	var u model.User
	err := collection.FindOneAndDelete(ctx, bson.M{"empid": empid}).Decode(&u)

	return &u, err
}

// // Update, will update user data by id
// // return error if any
// func Update(ctx context.Context, query interface{}, change interface{}) (*model.User, error) {
// 	conn := db.GetInstance()
// 	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
// 	var user model.User
// 	err := collection.FindOneAndUpdate(ctx, query, change).Decode(&user)
// 	return &user, err
// }

func UsersIsAlreadyExists(ctx context.Context, id string) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
	internalId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	var u model.User
	erro := collection.FindOne(ctx, bson.M{"_id": internalId}).Decode(&u)

	return erro == nil
}

func UsersIsAlreadyExistsWithEmpid(ctx context.Context, empid string) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)

	// @TODO: maybe validate empid argument with regex ?
	var u model.User
	erro := collection.FindOne(ctx, bson.M{"empid": empid}).Decode(&u)

	return erro == nil
}
