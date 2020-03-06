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

// UsersFindAll /
// UsersFindOne /
// UsersFindByEmpid /
// UsersFindOneById /
// UsersDeleteOne /
// UsersDeleteByEmpid /
// UsersUpdateOneById /
// UsersUpdateOne /
// UsersUpdateOneById /
// UsersUpdateOneByEmpid /
// UsersIsAlreadyExists /
// UsersIsAlreadyExistsWithId /
// UsersIsAlreadyExistsWithEmpid /

// UsersCollection name of users collection
const UsersCollection string = "users"

// UsersFindAll returns all users from repo
// @TODO pagination version of FindAll
func UsersFindAll(ctx context.Context) ([]*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)

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
	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
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
	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
	result, err := collection.InsertOne(ctx, u)

	if err != nil {
		return u, err
	}
	u.ID = result.InsertedID.(primitive.ObjectID)
	return u, nil
}

// UsersDeleteOne deletes a user and returns it
func UsersDeleteOne(ctx context.Context, query interface{}) (*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
	var u model.User
	err := collection.FindOneAndDelete(ctx, query).Decode(&u)
	return &u, err
}

// UsersDeleteByEmpid deletes a user by empid and returns it
func UsersDeleteByEmpid(ctx context.Context, empid string) (*model.User, error) {
	return UsersDeleteOne(ctx, bson.M{"empid": empid})
}

// UsersDeleteByEmpid deletes a user by empid and returns it
func UsersDeleteById(ctx context.Context, id string) (*model.User, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return UsersDeleteOne(ctx, bson.M{"_id": internalID})
}

// UsersDeleteOne deletes a user and returns it
// func UsersDeleteOne(ctx context.Context, user *model.User) (*model.User, error) {
// 	conn := db.GetInstance()
// 	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
// 	var u model.User
// 	err := collection.FindOneAndDelete(ctx, bson.M{"_id": user.ID}).Decode(&u)
// 	return &u, err
// }

// UsersDeleteByEmpid deletes a user by empid and returns it
// func UsersDeleteByEmpid(ctx context.Context, empid string) (*model.User, error) {
// 	conn := db.GetInstance()
// 	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
// 	var u model.User
// 	err := collection.FindOneAndDelete(ctx, bson.M{"empid": empid}).Decode(&u)

// 	return &u, err
// }

// UsersUpdateOne updates one
func UsersUpdateOne(ctx context.Context, query interface{}, update interface{}) (*model.User, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
	var u model.User
	err := collection.FindOneAndUpdate(ctx, query, update).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// UsersUpdateOneById updates one user by id
func UsersUpdateOneById(ctx context.Context, id string, update *model.User) (*model.User, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}

	return UsersUpdateOne(ctx, bson.M{"_id": internalID}, upd)
}

// UsersUpdateOneByEmpId updates a user by its empid
func UsersUpdateOneByEmpid(ctx context.Context, empid string, update *model.User) (*model.User, error) {
	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}
	return UsersUpdateOne(ctx, bson.M{"empid": empid}, upd)
}

// UsersIsAlreadyExists asks repo if user already exists
func UsersIsAlreadyExists(ctx context.Context, query interface{}) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
	var u model.User
	erro := collection.FindOne(ctx, query).Decode(&u)

	return erro == nil
}

// UsersIsAlreadyExists asks repo if user already exists, by id
func UsersIsAlreadyExistsWithId(ctx context.Context, id string) bool {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	return UsersIsAlreadyExists(ctx, bson.M{"_id": internalID})
}

// UsersIsAlreadyExistsWithEmpid asks repo if user already exists, by empid
func UsersIsAlreadyExistsWithEmpid(ctx context.Context, empid string) bool {
	return UsersIsAlreadyExists(ctx, bson.M{"empid": empid})
}
