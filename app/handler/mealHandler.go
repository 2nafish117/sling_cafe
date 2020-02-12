package handler

import (
	"context"
	"net/http"
	"sling_cafe/app/model"

	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/app/repo"
	"time"
)

// @TODO: streamline results

func MealPost(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	var meal model.Meal

	err := json.NewDecoder(request.Body).Decode(&meal)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	if !repo.UsersIsAlreadyExistsWithEmpid(context.TODO(), meal.EmpId) {
		// invalid empid
		// @TODO what status to return?
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte(`{ "message": "` + `empid: \"` + meal.EmpId + `\" doesnt exist` + `" }`))
		return
	}

	// register the time of meal
	meal.Timestamp = primitive.NewDateTimeFromTime(time.Now())
	// @TODO: automatically select breakfast/lunch/snacks ...

	m, err := repo.MealsInsertOne(context.TODO(), &meal)

	if err != nil {
		// @TODO: what status to return?
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message" : "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(m)

	// response.Header().Set("content-type", "application/json")
	// var meal model.Meal

	// err := json.NewDecoder(request.Body).Decode(&meal)

	// if err != nil {
	// 	response.WriteHeader(http.StatusBadRequest)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// conn := db.GetInstance()
	// users := conn.Database("test_db").Collection("users")

	// // TODO: check validity of emplotyee id
	// result := users.FindOne(context.TODO(), bson.M{"empid": meal.EmpId})
	// if err := result.Err(); err != nil {
	// 	// invalid empid
	// 	// @TODO what status to return?
	// 	response.WriteHeader(http.StatusUnauthorized)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }

	// meal.Timestamp = primitive.NewDateTimeFromTime(time.Now())
	// // TODO: automatically select breakfast/lunch/snacks ...

	// collection := conn.Database("test_db").Collection("meals")
	// insertResult, err := collection.InsertOne(context.TODO(), meal)

	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message" : "` + err.Error() + `" }`))
	// 	return
	// }

	// json.NewEncoder(response).Encode(insertResult)
}

func MealsGetByEmpid(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	empid, _ := params["empid"]

	meals, err := repo.MealsFindAllByEmpid(context.TODO(), empid)

	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(meals)

	// response.Header().Set("content-type", "application/json")

	// params := mux.Vars(request)
	// empid, _ := params["empid"]
	// conn := db.GetInstance()
	// collection := conn.Database("test_db").Collection("meals")
	// cursor, err := collection.Find(context.TODO(), bson.M{"empid": empid})
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// // defer cursor.Close(ctx)

	// if err := cursor.All(context.TODO(), &meals); err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// json.NewEncoder(response).Encode(meals)
}

// Only for debugging?
func MealsGet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	meals, err := repo.MealsFindAll(context.TODO())

	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(meals)

	// response.Header().Set("content-type", "application/json")
	// var meals []model.Meal
	// conn := db.GetInstance()
	// collection := conn.Database("test_db").Collection("meals")
	// cursor, err := collection.Find(context.TODO(), bson.M{})
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// // defer cursor.Close(ctx)

	// if err := cursor.All(context.TODO(), &meals); err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// json.NewEncoder(response).Encode(meals)
}
