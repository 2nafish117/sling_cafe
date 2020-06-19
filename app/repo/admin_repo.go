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

// AdminsCollection name of admins collection
const AdminsCollection string = "admins"

func AdminsFind(ctx context.Context, query interface{}) ([]*model.Admin, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(AdminsCollection)

	cursor, err := collection.Find(ctx, query)

	admins := make([]*model.Admin, 0)
	if err != nil {
		return admins, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &admins); err != nil {
		return admins, err
	}

	return admins, nil
}

func AdminsFindAll(ctx context.Context) ([]*model.Admin, error) {
	return AdminsFind(ctx, bson.M{})
}

func AdminsFindOne(ctx context.Context, query interface{}) (*model.Admin, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(AdminsCollection)
	var admin model.Admin
	err := collection.FindOne(ctx, query).Decode(&admin)
	return &admin, err
}

func AdminsFindByAdminID(ctx context.Context, adminID string) (*model.Admin, error) {
	return AdminsFindOne(ctx, bson.M{"admin_id": adminID})
}

func AdminsFindOneByID(ctx context.Context, id string) (*model.Admin, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return AdminsFindOne(ctx, bson.M{"_id": internalID})
}

func AdminsInsertOne(ctx context.Context, a *model.Admin) (*model.Admin, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(AdminsCollection)
	result, err := collection.InsertOne(ctx, a)

	if err != nil {
		return a, err
	}
	a.ID = result.InsertedID.(primitive.ObjectID)
	return a, nil
}

func AdminsDeleteOne(ctx context.Context, query interface{}) (*model.Admin, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(AdminsCollection)
	var a model.Admin
	err := collection.FindOneAndDelete(ctx, query).Decode(&a)
	return &a, err
}

func AdminsDeleteByAdminID(ctx context.Context, adminID string) (*model.Admin, error) {
	return AdminsDeleteOne(ctx, bson.M{"admin_id": adminID})
}

func AdminsDeleteByID(ctx context.Context, id string) (*model.Admin, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return AdminsDeleteOne(ctx, bson.M{"_id": internalID})
}

func AdminsUpdateOne(ctx context.Context, query interface{}, update interface{}) (*model.Admin, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(AdminsCollection)
	var a model.Admin
	err := collection.FindOneAndUpdate(ctx, query, update).Decode(&a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func AdminsUpdateOneByID(ctx context.Context, id string, update *model.Admin) (*model.Admin, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}

	return AdminsUpdateOne(ctx, bson.M{"_id": internalID}, upd)
}

func AdminsUpdateOneByAdminID(ctx context.Context, admin_id string, update *model.Admin) (*model.Admin, error) {
	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}
	return AdminsUpdateOne(ctx, bson.M{"admin_id": admin_id}, upd)
}

func AdminsIsAlreadyExists(ctx context.Context, query interface{}) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(AdminsCollection)
	var a model.Admin
	erro := collection.FindOne(ctx, query).Decode(&a)

	return erro == nil
}

func AdminsIsAlreadyExistsWithID(ctx context.Context, id string) bool {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	return AdminsIsAlreadyExists(ctx, bson.M{"_id": internalID})
}

func AdminsIsAlreadyExistsWithAdminID(ctx context.Context, adminID string) bool {
	return AdminsIsAlreadyExists(ctx, bson.M{"admin_id": adminID})
}
