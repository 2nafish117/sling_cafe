package handler

// import (
// 	"context"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"net/http"

// 	"sling_cafe/app/repo"
// 	"sling_cafe/util"
// 	"time"
// )

// // ReportsGet get bills of all users
// func ReportsGet(response http.ResponseWriter, request *http.Request) {

// 	queries := request.URL.Query()

// 	if queries.Get("all") == "true" || queries.Get("all") == "yes" {
// 		var fromDate primitive.DateTime = 0
// 		var toDate primitive.DateTime = 1<<63 - 1
// 		report, err := repo.ReportsFindOne(context.TODO(), fromDate, toDate)
// 		if err != nil {
// 			httpError := util.NewStatus(http.StatusNotFound, err.Error())
// 			util.Response(response, struct{}{}, httpError)
// 			return
// 		}
// 		util.Response(response, report, util.NewStatus(http.StatusOK, ""))
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

// 	report, err := repo.ReportsFindOne(context.TODO(), fromDate, toDate)
// 	if err != nil {
// 		httpError := util.NewStatus(http.StatusNotFound, err.Error())
// 		util.Response(response, struct{}{}, httpError)
// 		return
// 	}
// 	util.Response(response, report, util.NewStatus(http.StatusOK, ""))
// }
