package handler

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"

	"sling_cafe/app/repo"
	. "sling_cafe/log"
	"sling_cafe/util"
	"strconv"
	"time"
)

// ReceiptsGet get receipts of all users
func ReceiptsGet(response http.ResponseWriter, request *http.Request) {

	queries := request.URL.Query()

	// @TODO: add default values
	now := time.Now()
	oneMonthAgo := now.AddDate(0, 0, -30)

	// str := time.RFC3339Nano
	// layout := "2006-01-02T15:04:05Z"

	start, err := time.Parse("2006-01-02T15:04:05Z", util.GetOptQuery(queries.Get("start"), oneMonthAgo.Format("2006-01-02T15:04:05Z")))
	if err != nil {

		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	end, err := time.Parse("2006-01-02T15:04:05Z", util.GetOptQuery(queries.Get("end"), now.Format("2006-01-02T15:04:05Z")))
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	if start.UnixNano() > end.UnixNano() {
		httpError := util.NewStatus(http.StatusBadRequest, "start date is greater than end date")
		util.Response(response, struct{}{}, httpError)
		return
	}

	limit, err := strconv.Atoi(util.GetOptQuery(queries.Get("limit"), "500"))
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	pipeline := mongo.Pipeline{
		// Stage 1, match time between start_time and end_time
		bson.D{
			bson.E{
				Key: "$match", Value: bson.D{
					bson.E{Key: "time", Value: bson.D{
						bson.E{Key: "$gt", Value: primitive.NewDateTimeFromTime(start)},
						bson.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(end)},
					},
					},
				},
			},
		},
		// Stage 2, group by uid and add up the costs to find amtdue
		bson.D{bson.E{
			Key: "$group", Value: bson.D{
				bson.E{Key: "_id", Value: "$uid"},
				bson.E{Key: "amtdue", Value: bson.D{bson.E{Key: "$sum", Value: "$cost"}}},
			}}},

		// satge 3 sort based on amtdue
		bson.D{bson.E{
			Key: "$sort", Value: bson.D{
				bson.E{Key: "amtdue", Value: -1},
			}},
		},

		// stage 4 limit to limit
		bson.D{bson.E{
			Key: "$limit", Value: limit,
		},
		},
	}

	receipts, err := repo.ReceiptsAggregate(context.TODO(), pipeline)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, receipts, util.NewStatus(http.StatusOK, ""))
}

// ReceiptGetByUId gets an employees receipt
func ReceiptGetByUId(response http.ResponseWriter, request *http.Request) {
	queries := request.URL.Query()
	params := mux.Vars(request)
	Log.Info(queries)
	// @TODO: get time interval from query params
	// @TODO: what if startTime > endTime ???

	uid, _ := params["uid"]

	// @TODO: add default params

	// @TODO: assuming 30 day billing cycle, fix the billing cycle defaults based on range
	now := time.Now()
	oneMonthAgo := now.AddDate(0, 0, -30)

	start, err := time.Parse("2006-01-02T15:04:05Z", util.GetOptQuery(queries.Get("start"), oneMonthAgo.Format("2006-01-02T15:04:05Z")))
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	end, err := time.Parse("2006-01-02T15:04:05Z", util.GetOptQuery(queries.Get("end"), now.Format("2006-01-02T15:04:05Z")))
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	if start.UnixNano() > end.UnixNano() {
		httpError := util.NewStatus(http.StatusBadRequest, "start time is greater than end time")
		util.Response(response, struct{}{}, httpError)
		return
	}

	pipeline := mongo.Pipeline{
		// Stage 1, match time between start_time and end_time
		bson.D{
			bson.E{
				Key: "$match", Value: bson.D{
					bson.E{Key: "uid", Value: uid},
					bson.E{Key: "time", Value: bson.D{
						bson.E{Key: "$gt", Value: primitive.NewDateTimeFromTime(start)},
						bson.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(end)},
					},
					},
				},
			},
		},
		// Stage 2, group by uid and add up the costs to find amtdue
		bson.D{bson.E{
			Key: "$group", Value: bson.D{
				bson.E{Key: "_id", Value: "$uid"},
				bson.E{Key: "amtdue", Value: bson.D{bson.E{Key: "$sum", Value: "$cost"}}},
			}},
		},
	}

	receipts, err := repo.ReceiptsAggregate(context.TODO(), pipeline)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, receipts, util.NewStatus(http.StatusOK, ""))
}
