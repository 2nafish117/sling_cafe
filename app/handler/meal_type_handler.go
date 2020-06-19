package handler

import (
	"context"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sling_cafe/app/model"

	"encoding/json"

	"github.com/gorilla/mux"
	"sling_cafe/app/repo"
	"sling_cafe/util"
	// "time"
)

// time is stored and given as ISO time (with 0 difference from utc)
// application progem has to make sure of timezone conversions
// Internal mealType conversion adapter
// type adapterMealType struct {
// 	ID           primitive.ObjectID `json:"_id,omitempty"`
// 	MealID       string             `json:"meal_id,required"`
// 	Cost         int64              `json:"cost,required,string"`
// 	CompanyCost  int64              `json:"company_cost,required,string"`
// 	EmployeeCost int64              `json:"employee_cost,required,string"`
// 	FromTime     string             `json:"from_time,required"`
// 	ToTime       string             `json:"to_time,required"`
// 	// CatererID    string  `json:"caterer_id,required" bson:"caterer_id,required"`
// }

// func adapterToMealType(adapter *adapterMealType) (*model.MealType, error) {
// 	var mt model.MealType
// 	mt.MealID = adapter.MealID
// 	mt.Cost = adapter.Cost
// 	mt.CompanyCost = adapter.CompanyCost
// 	mt.EmployeeCost = adapter.EmployeeCost

// 	from, err := time.Parse(time.RFC3339, adapter.FromTime)
// 	if err != nil {
// 		return &mt, err
// 	}

// 	to, err := time.Parse(time.RFC3339, adapter.ToTime)
// 	if err != nil {
// 		return &mt, err
// 	}

// 	mt.FromTime = primitive.NewDateTimeFromTime(from)
// 	mt.ToTime = primitive.NewDateTimeFromTime(to)
// 	return &mt, nil
// }

// func mealTypeToAdapter(mt *model.MealType) *adapterMealType {
// 	var adapter adapterMealType
// 	adapter.ID = mt.ID
// 	adapter.MealID = mt.MealID
// 	adapter.Cost = mt.Cost
// 	adapter.CompanyCost = mt.CompanyCost
// 	adapter.EmployeeCost = mt.EmployeeCost
// 	adapter.FromTime = mt.FromTime.Time().Format(time.RFC3339)
// 	adapter.ToTime = mt.ToTime.Time().Format(time.RFC3339)

// 	return &adapter
// }

// MealTypePost posts a mealtype
func MealTypePost(response http.ResponseWriter, request *http.Request) {
	var mealType model.MealType

	err := json.NewDecoder(request.Body).Decode(&mealType)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	// var adapter adapterMealType

	// err := json.NewDecoder(request.Body).Decode(&adapter)
	// if err != nil {
	// 	httpError := util.NewStatus(http.StatusBadRequest, err.Error())
	// 	util.Response(response, struct{}{}, httpError)
	// 	return
	// }

	// // Time conversion stuff
	// mealType, err := adapterToMealType(&adapter)
	// if err != nil {
	// 	httpError := util.NewStatus(http.StatusBadRequest, err.Error())
	// 	util.Response(response, struct{}{}, httpError)
	// 	return
	// }

	if err := mealType.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if repo.MealTypesIsAlreadyExistsWithMealID(context.TODO(), mealType.MealID) {
		httpError := util.NewStatus(http.StatusBadRequest, "mealtype alredy exists, use update instead")
		util.Response(response, struct{}{}, httpError)
		return
	}

	mt, err := repo.MealTypesInsertOne(context.TODO(), &mealType)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	util.Response(response, mt, util.NewStatus(http.StatusOK, ""))

	// adapt := mealTypeToAdapter(mt)
	// util.Response(response, adapt, util.NewStatus(http.StatusOK, ""))
}

// MealTypesGet gets all mealtypes
func MealTypesGet(response http.ResponseWriter, request *http.Request) {
	mt, err := repo.MealTypesFindAll(context.TODO())
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	util.Response(response, mt, util.NewStatus(http.StatusOK, ""))

	// adapters := make([]*adapterMealType, 0, len(mt))
	// for _, it := range mt {
	// 	adapters = append(adapters, mealTypeToAdapter(it))
	// }
	// util.Response(response, adapters, util.NewStatus(http.StatusOK, ""))
}

// MealTypeGet gets a mealtype
func MealTypeGet(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	mealID, _ := params["meal_id"]

	mt, err := repo.MealTypesFindOneByMealID(context.TODO(), mealID)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}
	util.Response(response, mt, util.NewStatus(http.StatusOK, ""))

	// adapter := mealTypeToAdapter(mt)
	// util.Response(response, adapter, util.NewStatus(http.StatusOK, ""))
}

// MealTypePutByMealID puts by meal_id
func MealTypePutByMealID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	mealID, _ := params["meal_id"]

	var mealType model.MealType

	err := json.NewDecoder(request.Body).Decode(&mealType)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	// var adapter adapterMealType
	// err := json.NewDecoder(request.Body).Decode(&adapter)
	// if err != nil {
	// 	util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
	// 	return
	// }

	// mealType, err := adapterToMealType(&adapter)
	// if err != nil {
	// 	httpError := util.NewStatus(http.StatusBadRequest, err.Error())
	// 	util.Response(response, struct{}{}, httpError)
	// 	return
	// }

	if mealID != mealType.MealID {
		httpError := util.NewStatus(http.StatusBadRequest, "meal_id update and mealType meal_id mismatch")
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := mealType.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	mt, err := repo.MealTypesUpdateOneByMealID(context.TODO(), mealID, &mealType)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, mt, util.NewStatus(http.StatusOK, ""))

	// adapt := mealTypeToAdapter(mt)
	// util.Response(response, adapt, util.NewStatus(http.StatusOK, ""))
}

// MealTypeDeleteByMealID deletes user by uid
func MealTypeDeleteByMealID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	mealID, _ := params["meal_id"]

	mt, err := repo.MealTypesDeleteOneByMealID(context.TODO(), mealID)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, mt, util.NewStatus(http.StatusOK, ""))

	// adapter := mealTypeToAdapter(mt)
	// util.Response(response, adapter, util.NewStatus(http.StatusOK, ""))
}
