package handler

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sling_cafe/app/model"

	"encoding/json"
	"github.com/gorilla/mux"
	"sling_cafe/app/repo"
	"sling_cafe/util"
	"time"
)

// time is stored and given as ISO time (with 0 difference from utc)
// application progem has to make sure of timezone conversions
// Internal mealType conversion adapter
type adapterMealType struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type      string             `json:"type,omitempty" bson:"type,omitempty"`
	Cost      float64            `json:"cost,required" bson:"cost,required"`
	StartTime string             `json:"start_time,required" bson:"start_time,required"`
	EndTime   string             `json:"end_time,required" bson:"end_time,required"`
}

func adapterToMealType(adapter *adapterMealType) (*model.MealType, error) {
	var mt model.MealType
	mt.Type = adapter.Type
	mt.Cost = adapter.Cost

	// layout := "2006-01-02T15:04:05Z"
	start, err := time.Parse("2006-01-02T15:04:05Z", adapter.StartTime)
	if err != nil {
		return &mt, err
	}

	end, err := time.Parse("2006-01-02T15:04:05Z", adapter.EndTime)
	if err != nil {
		return &mt, err
	}

	mt.StartTime = primitive.NewDateTimeFromTime(start)
	mt.EndTime = primitive.NewDateTimeFromTime(end)
	return &mt, nil
}

func mealTypeToAdapter(mt *model.MealType) *adapterMealType {
	var adapter adapterMealType
	adapter.ID = mt.ID
	adapter.Type = mt.Type
	adapter.Cost = mt.Cost
	adapter.StartTime = mt.StartTime.Time().Format("2006-01-02T15:04:05Z")
	adapter.EndTime = mt.EndTime.Time().Format("2006-01-02T15:04:05Z")

	return &adapter
}

// MealTypePost posts a meal
func MealTypePost(response http.ResponseWriter, request *http.Request) {
	var adapter adapterMealType

	err := json.NewDecoder(request.Body).Decode(&adapter)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	// Time conversion stuff
	mealType, err := adapterToMealType(&adapter)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := mealType.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if repo.MealTypesIsAlreadyExistsWithType(context.TODO(), mealType.Type) {
		httpError := util.NewStatus(http.StatusBadRequest, "mealtype alredy exists, use update instead")
		util.Response(response, struct{}{}, httpError)
		return
	}

	mt, err := repo.MealTypesInsertOne(context.TODO(), mealType)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	// @TODO
	adapt := mealTypeToAdapter(mt)
	util.Response(response, adapt, util.NewStatus(http.StatusOK, ""))
}

// MealTypesGet gets meals eaten by uid
func MealTypesGet(response http.ResponseWriter, request *http.Request) {
	mt, err := repo.MealTypesFindAll(context.TODO())
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	adapters := make([]*adapterMealType, 0, len(mt))
	for _, it := range mt {
		adapters = append(adapters, mealTypeToAdapter(it))
	}
	util.Response(response, adapters, util.NewStatus(http.StatusOK, ""))
}

// MealTypeGet gets meals eaten by uid
func MealTypeGet(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	typ, _ := params["type"]

	mt, err := repo.MealTypesFindOneByType(context.TODO(), typ)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}
	adapter := mealTypeToAdapter(mt)
	util.Response(response, adapter, util.NewStatus(http.StatusOK, ""))
}

// MealTypePutByType gets meals eaten by uid
func MealTypePutByType(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	typ, _ := params["type"]

	var adapter adapterMealType
	err := json.NewDecoder(request.Body).Decode(&adapter)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	mealType, err := adapterToMealType(&adapter)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if typ != mealType.Type {
		httpError := util.NewStatus(http.StatusBadRequest, "type update and mealType type mismatch")
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := mealType.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	mt, err := repo.MealTypesUpdateOneByType(context.TODO(), typ, mealType)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	adapt := mealTypeToAdapter(mt)
	util.Response(response, adapt, util.NewStatus(http.StatusOK, ""))
}

// MealTypeDeleteByType deletes user by uid
func MealTypeDeleteByType(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	typ, _ := params["type"]

	mt, err := repo.MealTypesDeleteOneByType(context.TODO(), typ)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	adapter := mealTypeToAdapter(mt)
	util.Response(response, adapter, util.NewStatus(http.StatusOK, ""))
}
