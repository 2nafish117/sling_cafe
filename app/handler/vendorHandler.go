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

// VendorsGet gets all vendors
// @TODO get filters for pagination
func VendorsGet(response http.ResponseWriter, request *http.Request) {
	vendors, err := repo.VendorsFindAll(context.TODO())
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	util.Response(response, vendors, util.NewStatus(http.StatusOK, ""))
}

// VendorGet gets vendor
func VendorGet(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	id, _ := params["id"]
	vendor, err := repo.VendorsFindOneById(context.TODO(), id)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, vendor, util.NewStatus(http.StatusOK, ""))
}

// VendorGetByVid gets vendor by vid
func VendorGetByVid(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	vid, _ := params["vid"]

	vendor, err := repo.VendorsFindByVid(context.TODO(), vid)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, vendor, util.NewStatus(http.StatusOK, ""))
}

// VendorPost inserts a vendor
func VendorPost(response http.ResponseWriter, request *http.Request) {
	var vendor model.Vendor
	err := json.NewDecoder(request.Body).Decode(&vendor)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := vendor.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	// Check if vendor with vid already exists, if it does dont add that vendor
	if repo.VendorsIsAlreadyExistsWithVid(context.TODO(), vendor.VId) {
		httpError := util.NewStatus(http.StatusForbidden, "vendor already exists")
		util.Response(response, struct{}{}, httpError)
		return
	}

	v, err := repo.VendorsInsertOne(context.TODO(), &vendor)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, v, util.NewStatus(http.StatusOK, ""))
}

// VendorPutByVid updates vendor by id
func VendorPutByVid(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	vid, _ := params["vid"]

	var vendor model.Vendor
	err := json.NewDecoder(request.Body).Decode(&vendor)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := vendor.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	// Check if vendor with vid already exists, if it does dont add that vendor
	if !repo.VendorsIsAlreadyExistsWithVid(context.TODO(), vendor.VId) {
		httpError := util.NewStatus(http.StatusForbidden, "vendor doesnt exist to update!!")
		util.Response(response, struct{}{}, httpError)
		return
	}

	v, err := repo.VendorsUpdateOneByVid(context.TODO(), vid, &vendor)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, v, util.NewStatus(http.StatusOK, ""))
}

// VendorDelete deletes vendor by id
func VendorDelete(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	internalId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	vendor := &model.Vendor{ID: internalId}
	v, err := repo.VendorsDeleteOne(context.TODO(), vendor)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, v, util.NewStatus(http.StatusOK, ""))
}

// VendorDeleteByVid deletes vendor by vid
func VendorDeleteByVid(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	vid, _ := params["vid"]

	vendor, err := repo.VendorsDeleteByVid(context.TODO(), vid)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, vendor, util.NewStatus(http.StatusOK, ""))
}
