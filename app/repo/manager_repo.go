package repo

// import (
// 	"context"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	// "go.mongodb.org/mongo-driver/mongo/options"
// 	"mongo_test/db"
// 	"sling_cafe/app/model"
// 	"sling_cafe/config"
// 	"sling_cafe/util"
// )

// // AdminsFindAll /
// // AdminsFindOne /
// // AdminsFindByUId /
// // AdminsFindOneById /
// // AdminsDeleteOne /
// // AdminsDeleteByUId /
// // AdminsUpdateOneById /
// // AdminsUpdateOne /
// // AdminsUpdateOneById /
// // AdminsUpdateOneByUId /
// // AdminsIsAlreadyExists /
// // AdminsIsAlreadyExistsWithId /
// // AdminsIsAlreadyExistsWithUId /

// // UsersCollection name of users collection
// const AdminsCollection string = "admins"

// // UsersFindAll returns all users from repo
// // @TODO pagination version of FindAll
// // func UsersFindAll(ctx context.Context) ([]*model.User, error) {
// // 	conn := db.GetInstance()
// // 	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)

// // 	cursor, err := collection.Find(ctx, bson.M{})

// // 	var users []*model.User
// // 	if err != nil {
// // 		return users, err
// // 	}
// // 	defer cursor.Close(ctx)

// // 	if err := cursor.All(ctx, &users); err != nil {
// // 		return users, err
// // 	}

// // 	return users, nil
// // }

// // // UsersFindOne finds first user matching query and returns it
// // func UsersFindOne(ctx context.Context, query interface{}) (*model.User, error) {
// // 	conn := db.GetInstance()
// // 	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
// // 	var user model.User
// // 	err := collection.FindOne(ctx, query).Decode(&user)
// // 	return &user, err
// // }

// // // UsersFindByUId finds user with uid and returns
// // func UsersFindByUId(ctx context.Context, uid string) (*model.User, error) {
// // 	return UsersFindOne(ctx, bson.M{"uid": uid})
// // }

// // // UsersFindOneById find the user by id
// // func UsersFindOneById(ctx context.Context, id string) (*model.User, error) {
// // 	internalID, err := primitive.ObjectIDFromHex(id)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	return UsersFindOne(ctx, bson.M{"_id": internalID})
// // }

// // // UsersInsertOne inserts one user to the repository
// // func UsersInsertOne(ctx context.Context, u *model.User) (*model.User, error) {
// // 	conn := db.GetInstance()
// // 	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
// // 	result, err := collection.InsertOne(ctx, u)

// // 	if err != nil {
// // 		return u, err
// // 	}
// // 	u.ID = result.InsertedID.(primitive.ObjectID)
// // 	return u, nil
// // }

// // // UsersDeleteOne deletes a user and returns it
// // func UsersDeleteOne(ctx context.Context, query interface{}) (*model.User, error) {
// // 	conn := db.GetInstance()
// // 	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
// // 	var u model.User
// // 	err := collection.FindOneAndDelete(ctx, query).Decode(&u)
// // 	return &u, err
// // }

// // // UsersDeleteByUId deletes a user by uid and returns it
// // func UsersDeleteByUId(ctx context.Context, uid string) (*model.User, error) {
// // 	return UsersDeleteOne(ctx, bson.M{"uid": uid})
// // }

// // // UsersDeleteById deletes a user by id and returns it
// // func UsersDeleteById(ctx context.Context, id string) (*model.User, error) {
// // 	internalID, err := primitive.ObjectIDFromHex(id)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return UsersDeleteOne(ctx, bson.M{"_id": internalID})
// // }

// // // UsersUpdateOne updates one
// // func UsersUpdateOne(ctx context.Context, query interface{}, update interface{}) (*model.User, error) {
// // 	conn := db.GetInstance()
// // 	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
// // 	var u model.User
// // 	err := collection.FindOneAndUpdate(ctx, query, update).Decode(&u)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return &u, nil
// // }

// // // UsersUpdateOneById updates one user by id
// // func UsersUpdateOneById(ctx context.Context, id string, update *model.User) (*model.User, error) {
// // 	internalID, err := primitive.ObjectIDFromHex(id)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	customBson := util.CustomBson{}
// // 	upd, err := customBson.Set(update)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	return UsersUpdateOne(ctx, bson.M{"_id": internalID}, upd)
// // }

// // // UsersUpdateOneByUId updates a user by its uid
// // func UsersUpdateOneByUId(ctx context.Context, uid string, update *model.User) (*model.User, error) {
// // 	customBson := util.CustomBson{}
// // 	upd, err := customBson.Set(update)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return UsersUpdateOne(ctx, bson.M{"uid": uid}, upd)
// // }

// // // UsersIsAlreadyExists asks repo if user already exists
// // func UsersIsAlreadyExists(ctx context.Context, query interface{}) bool {
// // 	conn := db.GetInstance()
// // 	collection := conn.Database(config.GetInstance().DbName).Collection(UsersCollection)
// // 	var u model.User
// // 	erro := collection.FindOne(ctx, query).Decode(&u)

// // 	return erro == nil
// // }

// // // UsersIsAlreadyExistsWithId asks repo if user already exists, by id
// // func UsersIsAlreadyExistsWithId(ctx context.Context, id string) bool {
// // 	internalID, err := primitive.ObjectIDFromHex(id)
// // 	if err != nil {
// // 		return false
// // 	}
// // 	return UsersIsAlreadyExists(ctx, bson.M{"_id": internalID})
// // }

// // // UsersIsAlreadyExistsWithUId asks repo if user already exists, by uid
// // func UsersIsAlreadyExistsWithUId(ctx context.Context, uid string) bool {
// // 	return UsersIsAlreadyExists(ctx, bson.M{"uid": uid})
// // }
