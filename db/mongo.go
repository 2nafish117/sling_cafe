package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sling_cafe/config"
)

var dbInstance *mongo.Client = nil

// // init always runs only once, irresepective of number of times this package is imported
// func init() {

// 	// @TODO: context change?
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	cfg := config.GetInstance()

// 	clientOptions := options.Client().ApplyURI("mongodb://" + cfg.DbAddr)

// 	var err error
// 	dbInstance, err = mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Print(err.Error())

// 	}
// 	log.Print("connected to db at " + cfg.DbAddr)

// }

// call Connect manually from main, explicitly
func Connect() {

	cfg := config.GetInstance()

	clientOptions := options.Client().ApplyURI("mongodb://" + cfg.DbAddr)

	var err error
	dbInstance, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		// log.Print(err.Error())
		log.Panic(err.Error())
	}
	log.Print("connected to db at " + cfg.DbAddr)
}

// GetInstance Returns the singleton database instance
func GetInstance() *mongo.Client {
	return dbInstance
}
