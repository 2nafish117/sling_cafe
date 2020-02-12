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

// UsersFindAll returns all users from repo
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

// UsersFindOne finds first user matching query and returns it
func UsersFindOne(ctx context.Context, query interface{}) (*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
	var user model.User
	err := collection.FindOne(ctx, query).Decode(&user)
	return &user, err
}

// UsersFindByEmpid finds user with empid and returns
func UsersFindByEmpid(ctx context.Context, empid string) (*model.User, error) {
	return UsersFindOne(ctx, bson.M{"empid": empid})
}

// UsersFindOneById find the user by id
func UsersFindOneById(ctx context.Context, id string) (*model.User, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return UsersFindOne(ctx, bson.M{"_id": internalID})
}

// UsersInsertOne inserts one user to the repository
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

// UsersDeleteOne deletes a user and returns it
func UsersDeleteOne(ctx context.Context, user *model.User) (*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
	var u model.User
	err := collection.FindOneAndDelete(ctx, bson.M{"_id": user.ID}).Decode(&u)
	fmt.Print(u, err)
	return &u, err
}

// UsersDeleteByEmpid deletes a user by empid and returns it
func UsersDeleteByEmpid(ctx context.Context, empid string) (*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
	var u model.User
	err := collection.FindOneAndDelete(ctx, bson.M{"empid": empid}).Decode(&u)

	return &u, err
}

// UsersIsAlreadyExists asks repo if user already exists, by id
func UsersIsAlreadyExists(ctx context.Context, id string) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	var u model.User
	erro := collection.FindOne(ctx, bson.M{"_id": internalID}).Decode(&u)

	return erro == nil
}

// UsersIsAlreadyExistsWithEmpid asks repo if user already exists, by empid
func UsersIsAlreadyExistsWithEmpid(ctx context.Context, empid string) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(usersCollection)

	// @TODO: maybe validate empid argument with regex ?
	var u model.User
	erro := collection.FindOne(ctx, bson.M{"empid": empid}).Decode(&u)

	return erro == nil
}
