package handler

import (
	"context"
	"net/http"
	"sling_cafe/app/model"
	"sling_cafe/db"

	"encoding/json"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReceiptGet(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := params["id"]
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	var user model.User
	conn := db.GetInstance()
	collection := conn.Database("test_db").Collection("users")
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

func ReceiptsGet(response http.ResponseWriter, request *http.Request) {

}
