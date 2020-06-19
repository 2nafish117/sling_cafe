package handler

// import (
// 	"context"
// 	"github.com/gorilla/mux"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"net/http"
// 	"sling_cafe/app/model"

// 	"sling_cafe/app/repo"
// 	"sling_cafe/util"
// 	"time"
// )

// // BillsGet get bills of all users
// func BillsGet(response http.ResponseWriter, request *http.Request) {

// 	queries := request.URL.Query()

// 	sort := util.GetOptQuery(queries.Get("sort"), "d")
// 	var sortMode = -1
// 	if sort == "a" || sort == "asc" || sort == "ascending" {
// 		sortMode = 1
// 	}
// 	if queries.Get("all") == "true" || queries.Get("all") == "yes" {
// 		var fromDate primitive.DateTime = 0
// 		var toDate primitive.DateTime = 1<<63 - 1
// 		bills, err := repo.BillsFindAll(context.TODO(), fromDate, toDate, sortMode)
// 		if err != nil {
// 			httpError := util.NewStatus(http.StatusNotFound, err.Error())
// 			util.Response(response, struct{}{}, httpError)
// 			return
// 		}
// 		util.Response(response, bills, util.NewStatus(http.StatusOK, ""))
// 		return
// 	}

// 	start := queries.Get("start")
// 	end := queries.Get("end")

// 	startTime, err := time.Parse("2006-01-02T15:04:05Z", start)
// 	if err != nil {
// 		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
// 		util.Response(response, struct{}{}, httpError)
// 		return
// 	}
// 	endTime, err := time.Parse("2006-01-02T15:04:05Z", end)
// 	if err != nil {
// 		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
// 		util.Response(response, struct{}{}, httpError)
// 		return
// 	}

// 	if startTime.UnixNano() > endTime.UnixNano() {
// 		httpError := util.NewStatus(http.StatusBadRequest, "start date is greater than end date")
// 		util.Response(response, struct{}{}, httpError)
// 		return
// 	}

// 	fromDate := primitive.NewDateTimeFromTime(startTime)
// 	toDate := primitive.NewDateTimeFromTime(endTime)

// 	bills, err := repo.BillsFindAll(context.TODO(), fromDate, toDate, sortMode)
// 	if err != nil {
// 		httpError := util.NewStatus(http.StatusNotFound, err.Error())
// 		util.Response(response, struct{}{}, httpError)
// 		return
// 	}
// 	util.Response(response, bills, util.NewStatus(http.StatusOK, ""))
// }

// // BillGetByUId gets an users bill
// func BillGetByUId(response http.ResponseWriter, request *http.Request) {
// 	queries := request.URL.Query()
// 	params := mux.Vars(request)
// 	uid, _ := params["uid"]

// 	if !repo.UsersIsAlreadyExistsWithUId(context.TODO(), uid) {
// 		httpError := util.NewStatus(http.StatusNotFound, "user does not exist")
// 		util.Response(response, struct{}{}, httpError)
// 		return
// 	}

// 	if queries.Get("all") == "true" || queries.Get("all") == "yes" {
// 		var fromDate primitive.DateTime = 0
// 		var toDate primitive.DateTime = 1<<63 - 1
// 		bills, err := repo.BillsFindOne(context.TODO(), uid, fromDate, toDate)
// 		if err != nil {
// 			httpError := util.NewStatus(http.StatusNotFound, err.Error())
// 			util.Response(response, struct{}{}, httpError)
// 			return
// 		}
// 		util.Response(response, bills, util.NewStatus(http.StatusOK, ""))
// 		return
// 	}

// 	start := queries.Get("start")
// 	end := queries.Get("end")

// 	startTime, err := time.Parse("2006-01-02T15:04:05Z", start)
// 	if err != nil {
// 		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
// 		util.Response(response, struct{}{}, httpError)
// 		return
// 	}
// 	endTime, err := time.Parse("2006-01-02T15:04:05Z", end)
// 	if err != nil {
// 		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
// 		util.Response(response, struct{}{}, httpError)
// 		return
// 	}

// 	if startTime.UnixNano() > endTime.UnixNano() {
// 		httpError := util.NewStatus(http.StatusBadRequest, "start date is greater than end date")
// 		util.Response(response, struct{}{}, httpError)
// 		return
// 	}

// 	fromDate := primitive.NewDateTimeFromTime(startTime)
// 	toDate := primitive.NewDateTimeFromTime(endTime)

// 	bill, err := repo.BillsFindOne(context.TODO(), uid, fromDate, toDate)
// 	if err != nil {
// 		httpError := util.NewStatus(http.StatusNotFound, err.Error())
// 		util.Response(response, struct{}{}, httpError)
// 		return
// 	}

// 	if bill == nil {
// 		luckyUser, _ := repo.UsersFindByUId(context.TODO(), uid)
// 		emptyBill := model.Bill{
// 			User:      *luckyUser,
// 			Breakfast: model.MealBill{},
// 			Lunch:     model.MealBill{},
// 			Snack:     model.MealBill{},
// 			TotalDue:  0,
// 		}
// 		util.Response(response, emptyBill, util.NewStatus(http.StatusOK, ""))
// 		return
// 	}

// 	util.Response(response, bill, util.NewStatus(http.StatusOK, ""))
// }
