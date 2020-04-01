package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sling_cafe/config"
	. "sling_cafe/log"
)

var dbInstance *mongo.Client = nil

// call Connect manually from main, explicitly
func Connect() {

	cfg := config.GetInstance()

	clientOptions := options.Client().ApplyURI("mongodb://" + cfg.DbAddr)

	var err error
	dbInstance, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		Log.Fatalf(err.Error())
	}
	Log.Info("Connected to db at: ", cfg.DbAddr)
}

// GetInstance Returns the singleton database instance
func GetInstance() *mongo.Client {
	return dbInstance
}
