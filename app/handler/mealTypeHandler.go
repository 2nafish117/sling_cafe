package handler

import (
	"context"
	// "github.com/gorilla/mux"
	"net/http"
	"sling_cafe/app/model"

	"encoding/json"
	"sling_cafe/app/repo"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// MealTypePost handles POST to mealtypes
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

	mt, err := repo.MealTypesInsertOne(context.TODO(), &mealType)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message" : "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(mt)

	// response.Header().Set("content-type", "application/json")
	// var mealType model.MealType

	// // @TODO: DateTime internal representation vs evternal input??
	// err := json.NewDecoder(request.Body).Decode(&mealType)

	// if err != nil {
	// 	response.WriteHeader(http.StatusBadRequest)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }

	// conn := db.GetInstance()
	// collection := conn.Database("test_db").Collection("mealtypes")
	// result, err := collection.InsertOne(context.TODO(), mealType)

	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message" : "` + err.Error() + `" }`))
	// 	return
	// }

	// json.NewEncoder(response).Encode(result)
}

// MealTypesGet handles GET from mealtypes
func MealTypesGet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	mealTypes, err := repo.MealTypesFindAll(context.TODO())
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(mealTypes)

	// response.Header().Set("content-type", "application/json")
	// var mealTypes []model.MealType
	// conn := db.GetInstance()
	// collection := conn.Database("test_db").Collection("mealtypes")
	// cursor, err := collection.Find(context.TODO(), bson.M{})
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// // defer cursor.Close(ctx)

	// if err := cursor.All(context.TODO(), &mealTypes); err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// json.NewEncoder(response).Encode(mealTypes)
}

// @TODO: implement it
// MealTypesPut handles PUT to mealtypes
// func MealTypesPut(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("content-type", "application/json")
// 	params := mux.Vars(request)
// 	id, _ := params["mealtypeid"]

// 	var mealType model.MealType

// 	err := json.NewDecoder(request.Body).Decode(&mealType)
// 	if err != nil {
// 		response.WriteHeader(http.StatusBadRequest)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}

// 	mealTypes, err := repo.MealTypesUpdateOne(context.TODO(), &mealType)
// 	if err != nil {
// 		response.WriteHeader(http.StatusNotFound)
// 		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	json.NewEncoder(response).Encode(mealTypes)
// }
