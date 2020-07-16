package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "log"
	"net/http"
	"sling_cafe/app/model"
	"sling_cafe/app/repo"
	"sling_cafe/util"
	"time"
)

func MealEntryPost(response http.ResponseWriter, request *http.Request) {
	var meal model.MealEntry

	err := json.NewDecoder(request.Body).Decode(&meal)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := meal.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if !repo.EmployeesIsAlreadyExistsWithEmployeeID(context.TODO(), meal.EmployeeID) {
		// invalid employee_id
		httpError := util.NewStatus(http.StatusUnauthorized, "invalid employee_id")
		util.Response(response, struct{}{}, httpError)
		return
	}

	// register the time of meal
	now := time.Now()
	meal.DateTime = now

	// Automatically select breakfast/lunch/snacks ...
	mealType, err := repo.MealTypesFindOneByTimeOfDay(context.TODO(), now)
	if err != nil {
		// error in finding mealtype
		httpError := util.NewStatus(http.StatusForbidden, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if mealType.Inactive {
		httpError := util.NewStatus(http.StatusForbidden, "Meal Type is Inactive")
		util.Response(response, struct{}{}, httpError)
		return
	}

	if repo.CaterersIsInactive(context.TODO(), mealType.CatererID) {
		httpError := util.NewStatus(http.StatusForbidden, "Caterer is Inactive")
		util.Response(response, struct{}{}, httpError)
		return
	}

	meal.MealID = mealType.MealID

	// Get mealType using user given meal_id
	// mealType, err := repo.MealTypesFindOne(context.TODO(), bson.M{"meal_id": meal.MealID})
	// if err != nil {
	// 	// error in finding mealtype
	// 	httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
	// 	util.Response(response, struct{}{}, httpError)
	// 	return
	// }

	// Check if employee has already eaten the meal for that day
	if repo.MealEntriesHasEmployeeAlreadyEaten(context.TODO(), now, meal.EmployeeID, meal.MealID) {
		// already eaten today error
		httpError := util.NewStatus(http.StatusUnauthorized, "employee_id already eaten")
		util.Response(response, struct{}{}, httpError)
		return
	}

	meal.Cost = mealType.Cost
	meal.EmployeeCost = mealType.EmployeeCost
	meal.CompanyCost = mealType.CompanyCost
	meal.CatererID = mealType.CatererID
	m, err := repo.MealEntriesInsertOne(context.TODO(), &meal)

	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	util.Response(response, m, util.NewStatus(http.StatusOK, ""))
}

func MealEntriesGetByEmployeeID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	employeeID, _ := params["employee_id"]

	meals, err := repo.MealEntriesFindAllByEmployeeID(context.TODO(), employeeID)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	util.Response(response, meals, util.NewStatus(http.StatusOK, ""))
}

func MealEntriesGet(response http.ResponseWriter, request *http.Request) {
	meals, err := repo.MealEntriesFindAll(context.TODO())

	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	util.Response(response, meals, util.NewStatus(http.StatusOK, ""))
}
