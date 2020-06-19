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

// EmployeesCollection name of employee collection
const EmployeesCollection string = "employees"

func EmployeesFind(ctx context.Context, query interface{}) ([]*model.Employee, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(EmployeesCollection)

	cursor, err := collection.Find(ctx, query)

	employee := make([]*model.Employee, 0)
	if err != nil {
		return employee, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &employee); err != nil {
		return employee, err
	}

	return employee, nil
}

func EmployeesFindAll(ctx context.Context) ([]*model.Employee, error) {
	return EmployeesFind(ctx, bson.M{})
}

func EmployeesFindOne(ctx context.Context, query interface{}) (*model.Employee, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(EmployeesCollection)
	var Employee model.Employee
	err := collection.FindOne(ctx, query).Decode(&Employee)
	return &Employee, err
}

func EmployeesFindByEmployeeID(ctx context.Context, employeeID string) (*model.Employee, error) {
	return EmployeesFindOne(ctx, bson.M{"employee_id": employeeID})
}

func EmployeesFindOneByID(ctx context.Context, id string) (*model.Employee, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return EmployeesFindOne(ctx, bson.M{"_id": internalID})
}

func EmployeesInsertOne(ctx context.Context, u *model.Employee) (*model.Employee, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(EmployeesCollection)
	result, err := collection.InsertOne(ctx, u)

	if err != nil {
		return u, err
	}
	u.ID = result.InsertedID.(primitive.ObjectID)
	return u, nil
}

func EmployeesDeleteOne(ctx context.Context, query interface{}) (*model.Employee, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(EmployeesCollection)
	var u model.Employee
	err := collection.FindOneAndDelete(ctx, query).Decode(&u)
	return &u, err
}

func EmployeesDeleteByEmployeeID(ctx context.Context, employeeID string) (*model.Employee, error) {
	return EmployeesDeleteOne(ctx, bson.M{"employee_id": employeeID})
}

func EmployeesDeleteByID(ctx context.Context, id string) (*model.Employee, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return EmployeesDeleteOne(ctx, bson.M{"_id": internalID})
}

func EmployeesUpdateOne(ctx context.Context, query interface{}, update interface{}) (*model.Employee, error) {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(EmployeesCollection)
	var u model.Employee
	err := collection.FindOneAndUpdate(ctx, query, update).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func EmployeesUpdateOneByID(ctx context.Context, id string, update *model.Employee) (*model.Employee, error) {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}

	return EmployeesUpdateOne(ctx, bson.M{"_id": internalID}, upd)
}

func EmployeesUpdateOneByEmployeeID(ctx context.Context, employeeID string, update *model.Employee) (*model.Employee, error) {
	customBson := util.CustomBson{}
	upd, err := customBson.Set(update)
	if err != nil {
		return nil, err
	}
	return EmployeesUpdateOne(ctx, bson.M{"employee_id": employeeID}, upd)
}

func EmployeesIsAlreadyExists(ctx context.Context, query interface{}) bool {
	conn := db.GetInstance()
	collection := conn.Database(config.GetInstance().DbName).Collection(EmployeesCollection)
	var u model.Employee
	erro := collection.FindOne(ctx, query).Decode(&u)

	return erro == nil
}

func EmployeesIsAlreadyExistsWithID(ctx context.Context, id string) bool {
	internalID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	return EmployeesIsAlreadyExists(ctx, bson.M{"_id": internalID})
}

func EmployeesIsAlreadyExistsWithEmployeeID(ctx context.Context, employeeID string) bool {
	return EmployeesIsAlreadyExists(ctx, bson.M{"employee_id": employeeID})
}
