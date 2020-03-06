package handler

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"

	"errors"
	"sling_cafe/app/repo"
	. "sling_cafe/log"
	"sling_cafe/util"
	"time"
)

func ValidateDateFormat(start string, end string) (time.Time, time.Time, error) {
	startTime, errStart := time.Parse("2006-01-02", start)
	endTime, errEnd := time.Parse("2006-01-02", end)

	if errStart != nil {
		return startTime, endTime, errStart
	}

	if errEnd != nil {
		return startTime, endTime, errEnd
	}

	if startTime.UnixNano() > endTime.UnixNano() {
		return startTime, endTime, errors.New("start time is greater than end time")
	}

	return startTime, endTime, nil
}

// ReceiptsGet get receipts of all users
func ReceiptsGet(response http.ResponseWriter, request *http.Request) {

	queries := request.URL.Query()
	Log.Info(queries)
	// @TODO: get time interval from query params
	// @TODO: what date time precesion do i actually need ???
	startTime, endTime, err := ValidateDateFormat(queries.Get("start"), queries.Get("end"))
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusBadRequest, err.Error())
		util.Response(response, httpError)
		return
	}
	/*
		db.meals.aggregate([
			{ $match: { time: { $gt: ISODate("2020-01-26T03:47:42.213Z"), $lte:  ISODate("2020-04-05T07:34:35.304Z") } } },
			{ $group: { _id: "$empid", amtdue: { $sum: "$cost" } } }
		])
	*/

	pipeline := mongo.Pipeline{
		// Stage 1, match time between start_time and end_time
		bson.D{
			bson.E{
				Key: "$match", Value: bson.D{
					bson.E{Key: "time", Value: bson.D{
						bson.E{Key: "$gt", Value: primitive.NewDateTimeFromTime(startTime)},
						bson.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(endTime)},
					},
					},
				},
			},
		},
		// Stage 2, group by empid and add up the costs to find amtdue
		bson.D{bson.E{
			Key: "$group", Value: bson.D{
				bson.E{Key: "_id", Value: "$empid"},
				bson.E{Key: "amtdue", Value: bson.D{bson.E{Key: "$sum", Value: "$cost"}}},
			}},
		},
	}

	receipts, err := repo.ReceiptsAggregate(context.TODO(), pipeline)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusNotFound, err.Error())
		util.Response(response, httpError)
		return
	}
	util.Response(response, receipts)
}

// ReceiptGetById get receipt by _id
func ReceiptGetById(response http.ResponseWriter, request *http.Request) {
	Log.Info("not implemented")
	// queries := request.URL.Query()
	// params := mux.Vars(request)
	// Log.Info(queries)
	// // @TODO: get time interval from query params
	// // @TODO: what if startTime > endTime ???

	// id, _ := params["id"]
	// internalId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	httpError := util.NewErrorResponse(http.StatusBadRequest, err.Error())
	// 	util.Response(response, httpError)
	// 	return
	// }

	// startTime, errStart := time.Parse("2006-01-02", queries.Get("start"))
	// endTime, errEnd := time.Parse("2006-01-02", queries.Get("end"))

	// if errStart != nil {
	// 	httpError := util.NewErrorResponse(http.StatusBadRequest, "start time malformed")
	// 	util.Response(response, httpError)
	// 	return
	// }
	// if errEnd != nil {
	// 	httpError := util.NewErrorResponse(http.StatusBadRequest, "end time malformed")
	// 	util.Response(response, httpError)
	// 	return
	// }

	// if startTime.UnixNano() > endTime.UnixNano() {
	// 	httpError := util.NewErrorResponse(http.StatusBadRequest, "start time is greater than end time")
	// 	util.Response(response, httpError)
	// 	return
	// }

	// /*
	// 	db.meals.aggregate([
	// 		{ $match: { _id: "1" , time: { $gt: ISODate("2020-02-26T03:46:07.434Z"), $lte: ISODate("2020-03-05T07:25:28.972Z") } } },
	// 		{ $group: { _id: "$empid", amtdue: { $sum: "$cost" } } }
	// 	])
	// */

	// pipeline := mongo.Pipeline{
	// 	// Stage 1, match time between start_time and end_time
	// 	bson.D{
	// 		bson.E{
	// 			Key: "$match", Value: bson.D{
	// 				bson.E{Key: "_id", Value: internalId},
	// 				bson.E{Key: "time", Value: bson.D{
	// 					bson.E{Key: "$gt", Value: primitive.NewDateTimeFromTime(startTime)},
	// 					bson.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(endTime)},
	// 				},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	// Stage 2, group by _id and add up the costs to find amtdue
	// 	bson.D{bson.E{
	// 		Key: "$group", Value: bson.D{
	// 			bson.E{Key: "_id", Value: "$empid"}, // @TODO: what should _id of receipt be?
	// 			bson.E{Key: "amtdue", Value: bson.D{bson.E{Key: "$sum", Value: "$cost"}}},
	// 		}},
	// 	},
	// }

	// receipts, err := repo.ReceiptsAggregate(context.TODO(), pipeline)
	// if err != nil {
	// 	httpError := util.NewErrorResponse(http.StatusNotFound, err.Error())
	// 	util.Response(response, httpError)
	// 	return
	// }
	// util.Response(response, receipts)
}

// ReceiptGetByEmpid gets an employees receipt
func ReceiptGetByEmpid(response http.ResponseWriter, request *http.Request) {
	queries := request.URL.Query()
	params := mux.Vars(request)
	Log.Info(queries)
	// @TODO: get time interval from query params
	// @TODO: what if startTime > endTime ???

	empid, _ := params["empid"]

	// empid := queries.Get("empid")
	startTime, endTime, err := ValidateDateFormat(queries.Get("start"), queries.Get("end"))
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusBadRequest, err.Error())
		util.Response(response, httpError)
		return
	}

	/*
		db.meals.aggregate([
			{ $match: { empid: empid, time: { $gt: start_date, $lte: end_date } } },
			{ $group: { _id: "$empid", total: { $sum: "$cost" } } }
		])
	*/

	pipeline := mongo.Pipeline{
		// Stage 1, match time between start_time and end_time
		bson.D{
			bson.E{
				Key: "$match", Value: bson.D{
					bson.E{Key: "empid", Value: empid},
					bson.E{Key: "time", Value: bson.D{
						bson.E{Key: "$gt", Value: primitive.NewDateTimeFromTime(startTime)},
						bson.E{Key: "$lte", Value: primitive.NewDateTimeFromTime(endTime)},
					},
					},
				},
			},
		},
		// Stage 2, group by empid and add up the costs to find amtdue
		bson.D{bson.E{
			Key: "$group", Value: bson.D{
				bson.E{Key: "_id", Value: "$empid"},
				bson.E{Key: "amtdue", Value: bson.D{bson.E{Key: "$sum", Value: "$cost"}}},
			}},
		},
	}

	receipts, err := repo.ReceiptsAggregate(context.TODO(), pipeline)
	if err != nil {
		httpError := util.NewErrorResponse(http.StatusNotFound, err.Error())
		util.Response(response, httpError)
		return
	}
	util.Response(response, receipts)
}
