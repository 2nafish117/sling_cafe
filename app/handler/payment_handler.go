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
	"time"
)

// PaymentPost posts a payment
func PaymentPost(response http.ResponseWriter, request *http.Request) {
	var payment model.Payment

	err := json.NewDecoder(request.Body).Decode(&payment)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := payment.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if !repo.UsersIsAlreadyExistsWithUId(context.TODO(), payment.UId) {
		// invalid uid
		httpError := util.NewStatus(http.StatusUnauthorized, "invalid user")
		util.Response(response, struct{}{}, httpError)
		return
	}

	// register the time of payment
	payment.DateTime = primitive.NewDateTimeFromTime(time.Now())
	p, err := repo.PaymentsInsertOne(context.TODO(), &payment)

	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	util.Response(response, p, util.NewStatus(http.StatusOK, ""))
}

// PaymentsGetByUId gets payments eaten by uid
func PaymentsGetByUId(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	uid, _ := params["uid"]

	payments, err := repo.PaymentsFindAllByUId(context.TODO(), uid)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, payments, util.NewStatus(http.StatusOK, ""))
}

// PaymentsGet Only for debugging?
func PaymentsGet(response http.ResponseWriter, request *http.Request) {
	payments, err := repo.PaymentsFindAll(context.TODO())

	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, payments, util.NewStatus(http.StatusOK, ""))
}
