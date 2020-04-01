package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo_test/db"
	"sling_cafe/app/model"
	"sling_cafe/config"
	"sling_cafe/util"
)

// VendorsFindAll /
// VendorsFindOne /
// VendorsFindByuid /
// VendorsFindOneById /
// VendorsDeleteOne /
// VendorsDeleteByuid /
// VendorsUpdateOneById /
// VendorsUpdateOne /
// VendorsUpdateOneById /
// VendorsUpdateOneByuid /
// VendorsIsAlreadyExists /
// VendorsIsAlreadyExistsWithId /
// VendorsIsAlreadyExistsWithuid /

// VendorsCollection name of vendors collection
const VendorsCollection string = "vendors"

// VendorsFindAll returns all vendors from repo
// @TODO pagination version of FindAll
func VendorsFindAll(ctx context.Context) ([]*model.Vendor, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(VendorsCollection)

	cursor, err := collection.Find(ctx, bson.M{})

	var vendors []*model.Vendor
	if err != nil {
		return vendors, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &vendors); err != nil {
		return vendors, err
	}

	return vendors, nil
}

// VendorsFindOne finds first vendor matching query and returns it
func VendorsFindOne(ctx context.Context, query interface{}) (*model.Vendor, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(VendorsCollection)
	var vendor model.Vendor
	err := collection.FindOne(ctx, query).Decode(&vendor)
	return &vendor, err
}

// VendorsFindByVid finds vendor with empid and returns
func VendorsFindByVid(ctx context.Context, vid string) (*model.Vendor, error) {
	return VendorsFindOne(ctx, bson.M{"vid": vid})
}

// VendorsFindOneById find the vendor by id
func VendorsFindOneById(ctx context.Context, id string) (*model.Vendor, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return VendorsFindOne(ctx, bson.M{"_id": internalID})
}

// VendorsInsertOne inserts one vendor to the repository
func VendorsInsertOne(ctx context.Context, v *model.Vendor) (*model.Vendor, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(VendorsCollection)
	result, err := collection.InsertOne(ctx, v)

	if err != nil {
		return v, err
	}
	v.ID = result.InsertedID.(primitive.ObjectID)
	return v, nil
}

// VendorsDeleteOne deletes a vendor and returns it
func VendorsDeleteOne(ctx context.Context, query interface{}) (*model.Vendor, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(VendorsCollection)
	var v model.Vendor
	err := collection.FindOneAndDelete(ctx, query).Decode(&v)
	return &v, err
}

// VendorsDeleteByVid deletes a vendor by vid and returns it
func VendorsDeleteByVid(ctx context.Context, vid string) (*model.Vendor, error) {
	return VendorsDeleteOne(ctx, bson.M{"vid": vid})
}

// VendorsDeleteById deletes a vendor by id and returns it
func VendorsDeleteById(ctx context.Context, id string) (*model.Vendor, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return VendorsDeleteOne(ctx, bson.M{"_id": internalID})
}

// VendorsUpdateOne updates one
func VendorsUpdateOne(ctx context.Context, query interface{}, update interface{}) (*model.Vendor, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(VendorsCollection)
	var v model.Vendor
	err := collection.FindOneAndUpdate(ctx, query, update).Decode(&v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// VendorsUpdateOneById updates one vendor by id
func VendorsUpdateOneById(ctx context.Context, id string, update *model.Vendor) (*model.Vendor, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}

	return VendorsUpdateOne(ctx, bson.M{"_id": internalID}, upd)
}

// VendorsUpdateOneByVid updates a vendor by its vid
func VendorsUpdateOneByVid(ctx context.Context, vid string, update *model.Vendor) (*model.Vendor, error) {
	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}
	return VendorsUpdateOne(ctx, bson.M{"vid": vid}, upd)
}

// VendorsIsAlreadyExists asks repo if vendor already exists
func VendorsIsAlreadyExists(ctx context.Context, query interface{}) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(VendorsCollection)
	var v model.Vendor
	erro := collection.FindOne(ctx, query).Decode(&v)

	return erro == nil
}

// VendorsIsAlreadyExistsWithId asks repo if vendor already exists, by id
func VendorsIsAlreadyExistsWithId(ctx context.Context, id string) bool {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	return VendorsIsAlreadyExists(ctx, bson.M{"_id": internalID})
}

// VendorsIsAlreadyExistsWithVid asks repo if vendor already exists, by vid
func VendorsIsAlreadyExistsWithVid(ctx context.Context, vid string) bool {
	return VendorsIsAlreadyExists(ctx, bson.M{"vid": vid})
}
