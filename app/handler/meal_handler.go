package handler

import (
	"context"
	"net/http"
	"sling_cafe/app/model"

	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/app/repo"
	"sling_cafe/util"
	"time"
)

// time is stored and given as ISO time (with 0 difference from utc)
// application progem has to make sure of timezone conversions
// Internal mealType conversion adapter
type adapterMeal struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UId      string             `json:"uid,omitempty" bson:"uid,omitempty"`         // reference to an employee
	DateTime string             `json:"datetime,required" bson:"datetime,required"` // time is generated server side using time.Now()
	Type     string             `json:"type,required" bson:"type,required"`         // reference to a meal type
	Cost     float64            `json:"cost,required" bson:"cost,required"`
}

func mealToAdapter(m *model.Meal) *adapterMeal {
	var adapter adapterMeal
	adapter.ID = m.ID
	adapter.UId = m.UId
	adapter.DateTime = m.DateTime.Time().Format("2006-01-02T15:04:05Z")
	adapter.Type = m.Type
	adapter.Cost = m.Cost
	return &adapter
}

// MealPost posts a meal
func MealPost(response http.ResponseWriter, request *http.Request) {
	var meal model.Meal

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

	if !repo.UsersIsAlreadyExistsWithUId(context.TODO(), meal.UId) {
		// invalid uid
		httpError := util.NewStatus(http.StatusUnauthorized, "invalid user")
		util.Response(response, struct{}{}, httpError)
		return
	}

	// register the time of meal
	meal.DateTime = primitive.NewDateTimeFromTime(time.Now())
	// automatically select breakfast/lunch/snacks ...

	mealType, err := repo.MealtypesFindOne(context.TODO(), bson.M{"type": meal.Type})
	if err != nil {
		// error in finding mealtype
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	meal.Cost = mealType.Cost
	m, err := repo.MealsInsertOne(context.TODO(), &meal)

	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	util.Response(response, m, util.NewStatus(http.StatusOK, ""))
}

// MealsGetByUId gets meals eaten by uid
func MealsGetByUId(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	uid, _ := params["uid"]

	meals, err := repo.MealsFindAllByUId(context.TODO(), uid)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	adapters := make([]*adapterMeal, 0, len(meals))
	for _, it := range meals {
		adapters = append(adapters, mealToAdapter(it))
	}

	util.Response(response, adapters, util.NewStatus(http.StatusOK, ""))
}

// MealsGet Only for debugging?
func MealsGet(response http.ResponseWriter, request *http.Request) {
	meals, err := repo.MealsFindAll(context.TODO())

	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	adapters := make([]*adapterMeal, 0, len(meals))
	for _, it := range meals {
		adapters = append(adapters, mealToAdapter(it))
	}

	util.Response(response, adapters, util.NewStatus(http.StatusOK, ""))
}
