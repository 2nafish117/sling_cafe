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

func AdminsGet(response http.ResponseWriter, request *http.Request) {
	admins, err := repo.AdminsFindAll(context.TODO())
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}

	util.Response(response, admins, util.NewStatus(http.StatusOK, ""))
}

func AdminGet(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	id, _ := params["id"]
	admin, err := repo.AdminsFindOneByID(context.TODO(), id)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}
	util.Response(response, admin, util.NewStatus(http.StatusOK, ""))
}

func AdminGetByAdminID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	adminID, _ := params["admin_id"]

	admin, err := repo.AdminsFindByAdminID(context.TODO(), adminID)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}
	util.Response(response, admin, util.NewStatus(http.StatusOK, ""))
}

func AdminPost(response http.ResponseWriter, request *http.Request) {
	var admin model.Admin
	err := json.NewDecoder(request.Body).Decode(&admin)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	if err := admin.Validate(); err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	// Check if admin with admin_id already exists, if it does dont add that admin
	if repo.AdminsIsAlreadyExistsWithAdminID(context.TODO(), admin.AdminID) {
		httpError := util.NewStatus(http.StatusForbidden, "admin already exists")
		util.Response(response, struct{}{}, httpError)
		return
	}

	e, err := repo.AdminsInsertOne(context.TODO(), &admin)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusInternalServerError, err.Error()))
		return
	}
	util.Response(response, e, util.NewStatus(http.StatusOK, ""))
}

func AdminPutByAdminID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	admin_id, _ := params["admin_id"]

	var admin model.Admin
	err := json.NewDecoder(request.Body).Decode(&admin)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	if admin_id != admin.AdminID {
		httpError := util.NewStatus(http.StatusBadRequest, "admin_id update and admin admin_id mismatch")
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := admin.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	// Check if admin with admin_id already exists, if it does update
	if !repo.AdminsIsAlreadyExistsWithAdminID(context.TODO(), admin.AdminID) {
		httpError := util.NewStatus(http.StatusForbidden, "admin doesnt exist to update")
		util.Response(response, struct{}{}, httpError)
		return
	}

	e, err := repo.AdminsUpdateOneByAdminID(context.TODO(), admin_id, &admin)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, e, util.NewStatus(http.StatusOK, ""))
}

func AdminDelete(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	internalId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	admin := &model.Admin{ID: internalId}
	e, err := repo.AdminsDeleteOne(context.TODO(), admin)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, e, util.NewStatus(http.StatusOK, ""))
}

func AdminDeleteByAdminID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	admin_id, _ := params["admin_id"]

	admin, err := repo.AdminsDeleteByAdminID(context.TODO(), admin_id)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, admin, util.NewStatus(http.StatusOK, ""))
}
