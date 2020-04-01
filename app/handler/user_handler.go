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

// UsersGet gets all users
// @TODO get filters for pagination
func UsersGet(response http.ResponseWriter, request *http.Request) {
	users, err := repo.UsersFindAll(context.TODO())
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}

	util.Response(response, users, util.NewStatus(http.StatusOK, ""))
}

// UserGet gets user
func UserGet(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	id, _ := params["id"]
	user, err := repo.UsersFindOneById(context.TODO(), id)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}
	util.Response(response, user, util.NewStatus(http.StatusOK, ""))
}

// UserGetByUId gets user by uid
func UserGetByUId(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	uid, _ := params["uid"]

	user, err := repo.UsersFindByUId(context.TODO(), uid)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}
	util.Response(response, user, util.NewStatus(http.StatusOK, ""))
}

// UserPost inserts a user
func UserPost(response http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	if err := user.Validate(); err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	// Check if user with uid already exists, if it does dont add that user
	if repo.UsersIsAlreadyExistsWithUId(context.TODO(), user.UId) {
		httpError := util.NewStatus(http.StatusForbidden, "user already exists")
		util.Response(response, struct{}{}, httpError)
		return
	}

	u, err := repo.UsersInsertOne(context.TODO(), &user)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusInternalServerError, err.Error()))
		return
	}
	util.Response(response, u, util.NewStatus(http.StatusOK, err.Error()))
}

// UserPutByUId updates user by id
func UserPutByUId(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	uid, _ := params["uid"]

	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	if uid != user.UId {
		httpError := util.NewStatus(http.StatusBadRequest, "uid update and user uid mismatch")
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := user.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	// Check if user with uid already exists, if it does update
	if !repo.UsersIsAlreadyExistsWithUId(context.TODO(), user.UId) {
		httpError := util.NewStatus(http.StatusForbidden, "user doesnt exist to update")
		util.Response(response, struct{}{}, httpError)
		return
	}

	u, err := repo.UsersUpdateOneByUId(context.TODO(), uid, &user)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, u, util.NewStatus(http.StatusOK, ""))
}

// UserDelete deletes user by id
func UserDelete(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	internalId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	user := &model.User{ID: internalId}
	u, err := repo.UsersDeleteOne(context.TODO(), user)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, u, util.NewStatus(http.StatusOK, ""))
}

// UserDeleteByUId deletes user by uid
func UserDeleteByUId(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	uid, _ := params["uid"]

	user, err := repo.UsersDeleteByUId(context.TODO(), uid)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, user, util.NewStatus(http.StatusOK, ""))
}
