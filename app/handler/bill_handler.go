package handler

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"sling_cafe/app/repo"
	"sling_cafe/util"
	"time"
)

// BillsGet get receipts of all users
func BillsGet(response http.ResponseWriter, request *http.Request) {

	queries := request.URL.Query()

	start := queries.Get("start")
	end := queries.Get("end")
	sort := util.GetOptQuery(queries.Get("sort"), "d")
	var sortMode = -1
	if sort == "a" {
		sortMode = 1
	}

	startTime, err := time.Parse("2006-01-02T15:04:05Z", start)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	endTime, err := time.Parse("2006-01-02T15:04:05Z", end)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if startTime.UnixNano() > endTime.UnixNano() {
		httpError := util.NewStatus(http.StatusBadRequest, "start date is greater than end date")
		util.Response(response, struct{}{}, httpError)
		return
	}

	fromDate := primitive.NewDateTimeFromTime(startTime)
	toDate := primitive.NewDateTimeFromTime(endTime)

	bills, err := repo.BillsFindAll(context.TODO(), fromDate, toDate, sortMode)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, bills, util.NewStatus(http.StatusOK, ""))
}

// BillGetByUId gets an users bill
func BillGetByUId(response http.ResponseWriter, request *http.Request) {
	queries := request.URL.Query()
	params := mux.Vars(request)
	uid, _ := params["uid"]

	start := queries.Get("start")
	end := queries.Get("end")

	startTime, err := time.Parse("2006-01-02T15:04:05Z", start)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	endTime, err := time.Parse("2006-01-02T15:04:05Z", end)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if startTime.UnixNano() > endTime.UnixNano() {
		httpError := util.NewStatus(http.StatusBadRequest, "start date is greater than end date")
		util.Response(response, struct{}{}, httpError)
		return
	}

	fromDate := primitive.NewDateTimeFromTime(startTime)
	toDate := primitive.NewDateTimeFromTime(endTime)

	if !repo.UsersIsAlreadyExistsWithUId(context.TODO(), uid) {
		httpError := util.NewStatus(http.StatusNotFound, "user does not exist")
		util.Response(response, struct{}{}, httpError)
		return
	}

	bill, err := repo.BillsFindOne(context.TODO(), uid, fromDate, toDate)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, bill, util.NewStatus(http.StatusOK, ""))
}
