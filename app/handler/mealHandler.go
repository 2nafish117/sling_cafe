package handler

import (
	"context"
	"net/http"
	"sling_cafe/app/model"

	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/app/repo"
	"sling_cafe/util"
	"time"
)

// @TODO: streamline results

// MealPost posts a meal
func MealPost(response http.ResponseWriter, request *http.Request) {

	// response.Header().Set("content-type", "application/json")
	var meal model.Meal

	err := json.NewDecoder(request.Body).Decode(&meal)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusBadRequest, err.Error())
		util.Response(response, httpError)
		return
		// response.WriteHeader(http.StatusBadRequest)
		// response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	}

	if err := meal.Validate(); err != nil {
		httpError := util.NewErrorResponse(http.StatusBadRequest, err.Error())
		util.Response(response, httpError)
		return
	}

	if !repo.UsersIsAlreadyExistsWithEmpid(context.TODO(), meal.EmpId) {
		// invalid empid
		// @TODO what status to return?
		httpError := util.NewErrorResponse(http.StatusUnauthorized, "invalid empid")
		util.Response(response, httpError)
		// response.WriteHeader(http.StatusUnauthorized)
		// response.Write([]byte(`{ "message": "` + `empid: \"` + meal.EmpId + `\" doesnt exist` + `" }`))
		return
	}

	// register the time of meal
	meal.Timestamp = primitive.NewDateTimeFromTime(time.Now())
	// @TODO: automatically select breakfast/lunch/snacks ...

	m, err := repo.MealsInsertOne(context.TODO(), &meal)

	if err != nil {
		// @TODO: what status to return?
		httpError := util.NewErrorResponse(http.StatusInternalServerError, err.Error())
		util.Response(response, httpError)
		// response.WriteHeader(http.StatusInternalServerError)
		// response.Write([]byte(`{ "message" : "` + err.Error() + `" }`))
		return
	}

	util.Response(response, m)
	// json.NewEncoder(response).Encode(m)
}

// MealsGetByEmpid gets meals eaten by empid
func MealsGetByEmpid(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	empid, _ := params["empid"]

	meals, err := repo.MealsFindAllByEmpid(context.TODO(), empid)

	if err != nil {
		httpError := util.NewErrorResponse(http.StatusNotFound, err.Error())
		util.Response(response, httpError)
		// response.WriteHeader(http.StatusNotFound)
		// response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	util.Response(response, meals)
	// json.NewEncoder(response).Encode(meals)
}

// MealsGet Only for debugging?
func MealsGet(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("content-type", "application/json")

	meals, err := repo.MealsFindAll(context.TODO())

	if err != nil {
		httpError := util.NewErrorResponse(http.StatusNotFound, err.Error())
		util.Response(response, httpError)
		// response.WriteHeader(http.StatusNotFound)
		// response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	util.Response(response, meals)
}
