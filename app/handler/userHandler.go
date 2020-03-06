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
// @TODO: return mongo db id or empid?
func UsersGet(response http.ResponseWriter, request *http.Request) {
	users, err := repo.UsersFindAll(context.TODO())
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusNotFound, err.Error())
		util.Response(response, httpError)
		return
	}

	util.Response(response, users)
}

// UserGet gets user
func UserGet(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	id, _ := params["id"]
	user, err := repo.UsersFindOneById(context.TODO(), id)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusNotFound, err.Error())
		util.Response(response, httpError)
		// response.WriteHeader(http.StatusNotFound)
		// response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	util.Response(response, user)
	// json.NewEncoder(response).Encode(user)
}

// UserGetByEmpid gets user by empid
func UserGetByEmpid(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	empid, _ := params["empid"]

	user, err := repo.UsersFindByEmpid(context.TODO(), empid)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusNotFound, err.Error())
		util.Response(response, httpError)
		return
	}
	util.Response(response, user)
}

// UserPost inserts a user
func UserPost(response http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusBadRequest, err.Error())
		util.Response(response, httpError)
		return
	}

	if err := user.Validate(); err != nil {
		httpError := util.NewErrorResponse(http.StatusBadRequest, err.Error())
		util.Response(response, httpError)
		return
	}

	// Check if user with empid already exists, if it does dont add that user
	if repo.UsersIsAlreadyExistsWithEmpid(context.TODO(), user.EmpId) {
		httpError := util.NewErrorResponse(http.StatusForbidden, "user already exists")
		util.Response(response, httpError)
		return
	}

	u, err := repo.UsersInsertOne(context.TODO(), &user)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusInternalServerError, err.Error())
		util.Response(response, httpError)
		return
	}
	util.Response(response, u)
}

// UserPutByEmpid updates user by id
func UserPutByEmpid(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	empid, _ := params["empid"]

	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusBadRequest, err.Error())
		util.Response(response, httpError)
		return
	}

	if err := user.Validate(); err != nil {
		httpError := util.NewErrorResponse(http.StatusBadRequest, err.Error())
		util.Response(response, httpError)
		return
	}

	// Check if user with empid already exists, if it does dont add that user
	if !repo.UsersIsAlreadyExistsWithEmpid(context.TODO(), user.EmpId) {
		httpError := util.NewErrorResponse(http.StatusForbidden, "user doesnt exist to update!!")
		util.Response(response, httpError)
		return
	}

	u, err := repo.UsersUpdateOneByEmpid(context.TODO(), empid, &user)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusInternalServerError, err.Error())
		util.Response(response, httpError)
		return
	}
	util.Response(response, u)
}

// UserDelete deletes user by id
func UserDelete(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	internalId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusBadRequest, err.Error())
		util.Response(response, httpError)
		return
	}
	user := &model.User{ID: internalId}
	u, err := repo.UsersDeleteOne(context.TODO(), user)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusNotFound, err.Error())
		util.Response(response, httpError)
		return
	}
	util.Response(response, u)
}

// UserDeleteByEmpid deletes user by empid
func UserDeleteByEmpid(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	empid, _ := params["empid"]

	user, err := repo.UsersDeleteByEmpid(context.TODO(), empid)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusNotFound, err.Error())
		util.Response(response, httpError)
		return
	}
	util.Response(response, user)
}
