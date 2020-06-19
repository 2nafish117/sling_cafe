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
)

func CaterersGet(response http.ResponseWriter, request *http.Request) {
	employees, err := repo.CaterersFindAll(context.TODO())
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}

	util.Response(response, employees, util.NewStatus(http.StatusOK, ""))
}

func CatererGet(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	id, _ := params["id"]
	caterer, err := repo.CaterersFindOneByID(context.TODO(), id)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}
	util.Response(response, caterer, util.NewStatus(http.StatusOK, ""))
}

func CatererGetByCatererID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	employeeID, _ := params["caterer_id"]

	caterer, err := repo.CaterersFindByCatererID(context.TODO(), employeeID)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}
	util.Response(response, caterer, util.NewStatus(http.StatusOK, ""))
}

func CatererPost(response http.ResponseWriter, request *http.Request) {
	var caterer model.Caterer
	err := json.NewDecoder(request.Body).Decode(&caterer)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	if err := caterer.Validate(); err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	// Check if caterer with caterer_id already exists, if it does dont add that caterer
	if repo.CaterersIsAlreadyExistsWithCatererID(context.TODO(), caterer.CatererID) {
		httpError := util.NewStatus(http.StatusForbidden, "caterer already exists")
		util.Response(response, struct{}{}, httpError)
		return
	}

	e, err := repo.CaterersInsertOne(context.TODO(), &caterer)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusInternalServerError, err.Error()))
		return
	}
	util.Response(response, e, util.NewStatus(http.StatusOK, ""))
}

func CatererPutByCatererID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	caterer_id, _ := params["caterer_id"]

	var caterer model.Caterer
	err := json.NewDecoder(request.Body).Decode(&caterer)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	if caterer_id != caterer.CatererID {
		httpError := util.NewStatus(http.StatusBadRequest, "caterer_id update and caterer caterer_id mismatch")
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := caterer.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	// Check if caterer with caterer_id already exists, if it does update
	if !repo.CaterersIsAlreadyExistsWithCatererID(context.TODO(), caterer.CatererID) {
		httpError := util.NewStatus(http.StatusForbidden, "caterer doesnt exist to update")
		util.Response(response, struct{}{}, httpError)
		return
	}

	e, err := repo.CaterersUpdateOneByCatererID(context.TODO(), caterer_id, &caterer)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, e, util.NewStatus(http.StatusOK, ""))
}

func CatererDelete(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	internalId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	caterer := &model.Caterer{ID: internalId}
	e, err := repo.CaterersDeleteOne(context.TODO(), caterer)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, e, util.NewStatus(http.StatusOK, ""))
}

func CatererDeleteByCatererID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	caterer_id, _ := params["caterer_id"]

	caterer, err := repo.CaterersDeleteByCatererID(context.TODO(), caterer_id)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, caterer, util.NewStatus(http.StatusOK, ""))
}
