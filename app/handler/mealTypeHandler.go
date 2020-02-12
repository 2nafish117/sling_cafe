package handler

import (
	"context"
	"net/http"
	"sling_cafe/app/model"
	"sling_cafe/db"

	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func MealTypePost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var mealType model.MealType

	// @TODO: DateTime internal representation vs evternal input??
	err := json.NewDecoder(request.Body).Decode(&mealType)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	conn := db.GetInstance()
	collection := conn.Database("test_db").Collection("mealtypes")
	result, err := collection.InsertOne(context.TODO(), mealType)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message" : "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(result)
}

func MealTypesGet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var mealTypes []model.MealType
	conn := db.GetInstance()
	collection := conn.Database("test_db").Collection("mealtypes")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	// defer cursor.Close(ctx)

	if err := cursor.All(context.TODO(), &mealTypes); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(mealTypes)
}
