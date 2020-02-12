package handler

import (
	"context"
	"net/http"
	"sling_cafe/app/model"

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/app/repo"
)

// @TODO get filters for pagination
// @TODO: return mongo db id or empid?
func UsersGet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	users, err := repo.UsersFindAll(context.TODO())
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// @TODO: streamline responses
	type UsersResponse struct {
		U []*model.User `json:"result,required" bson:"result,required"`
	}
	ur := UsersResponse{U: users}
	fmt.Print(ur)
	json.NewEncoder(response).Encode(ur)

	// var users []model.User
	// conn := db.GetInstance()
	// collection := conn.Database("test_db").Collection("users")
	// cursor, err := collection.Find(context.TODO(), bson.M{})
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// // defer cursor.Close(ctx)

	// if err := cursor.All(context.TODO(), &users); err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// json.NewEncoder(response).Encode(users)
}

// UserGet gets user
func UserGet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	id, _ := params["id"]
	user, err := repo.UsersFindOneById(context.TODO(), id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)

	// id, err := primitive.ObjectIDFromHex(params["id"])

	// if err != nil {
	// 	response.WriteHeader(http.StatusBadRequest)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }

	// var user model.User
	// conn := db.GetInstance()
	// collection := conn.Database("test_db").Collection("users")
	// err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// json.NewEncoder(response).Encode(user)
}

func UserGetByEmpid(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	empid, _ := params["empid"]

	user, err := repo.UsersFindByEmpid(context.TODO(), empid)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)

	// var user model.User
	// conn := db.GetInstance()
	// collection := conn.Database("test_db").Collection("users")
	// err := collection.FindOne(context.TODO(), bson.M{"empid": empid}).Decode(&user)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// // defer cursor.Close(ctx)

	// json.NewEncoder(response).Encode(user)
}

func UserPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// Check if user with empid already exists, if it does dont add that user
	if repo.UsersIsAlreadyExistsWithEmpid(context.TODO(), user.EmpId) {
		response.WriteHeader(http.StatusForbidden)
		response.Write([]byte(`{ "message" : "` + `user with empid: \"` + user.EmpId + `\" already exists` + `" }`))
		return
	}

	u, err := repo.UsersInsertOne(context.TODO(), &user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message" : "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(model.UserResponse{Empid: u.EmpId})

	// var user model.User

	// err := json.NewDecoder(request.Body).Decode(&user)

	// if err != nil {
	// 	response.WriteHeader(http.StatusBadRequest)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }

	// conn := db.GetInstance()
	// collection := conn.Database("test_db").Collection("users")
	// result, err := collection.InsertOne(context.TODO(), user)

	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message" : "` + err.Error() + `" }`))
	// 	return
	// }

	// json.NewEncoder(response).Encode(result)
}

func UserDelete(response http.ResponseWriter, request *http.Request) {
	fmt.Print("UserDelete")
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	internalId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	user := &model.User{ID: internalId}
	u, err := repo.UsersDeleteOne(context.TODO(), user)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(u)

	// var user model.User
	// conn := db.GetInstance()
	// collection := conn.Database("test_db").Collection("users")
	// err = collection.FindOneAndDelete(context.TODO(), bson.M{"_id": id}).Decode(&user)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// // defer cursor.Close(ctx)

	// json.NewEncoder(response).Encode(user)
}

func UserDeleteByEmpid(response http.ResponseWriter, request *http.Request) {
	fmt.Print("UserDeleteByEmpid")
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	empid, _ := params["empid"]

	user, err := repo.UsersDeleteByEmpid(context.TODO(), empid)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)

	// var user model.User
	// conn := db.GetInstance()
	// collection := conn.Database("test_db").Collection("users")
	// err := collection.FindOneAndDelete(context.TODO(), bson.M{"empid": empid}).Decode(&user)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }

	// json.NewEncoder(response).Encode(user)
}
