package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sling_cafe/app/model"
	"sling_cafe/app/repo"
	"sling_cafe/util"
	"time"
)

// // time is stored and given as ISO time (with 0 difference from utc)
// // application progem has to make sure of timezone conversions
// // Internal mealType conversion adapter
// type adapterPayment struct {
// 	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	UId      string             `json:"uid,omitempty" bson:"uid,omitempty"`         // reference to an employee
// 	DateTime string             `json:"datetime,required" bson:"datetime,required"` // time is generated server side using time.Now()
// 	Mode     string             `json:"mode,required" bson:"mode,required"`         // mode of transaction
// 	Amount   float64            `json:"amount,required" bson:"amount,required"`
// }

// func paymentToAdapter(m *model.AdminTransaction) *adapterPayment {
// 	var adapter adapterPayment
// 	adapter.ID = m.ID
// 	adapter.UId = m.AdminID
// 	adapter.DateTime = m.DateTime.Time().Format("2006-01-02T15:04:05Z")
// 	adapter.Mode = m.Mode
// 	adapter.Amount = m.Amount
// 	return &adapter
// }

func AdminTransactionPost(response http.ResponseWriter, request *http.Request) {
	var transaction model.AdminTransaction

	err := json.NewDecoder(request.Body).Decode(&transaction)
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := transaction.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	if !repo.AdminsIsAlreadyExistsWithAdminID(context.TODO(), transaction.AdminID) {
		// invalid uid
		httpError := util.NewStatus(http.StatusUnauthorized, "invalid admin_id")
		util.Response(response, struct{}{}, httpError)
		return
	}

	// register the time of transaction
	transaction.DateTime = time.Now()
	p, err := repo.AdminTransactionsInsertOne(context.TODO(), &transaction)

	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	util.Response(response, p, util.NewStatus(http.StatusOK, ""))
}

func AdminTransactionsGetByAdminID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	adminID, _ := params["admin_id"]

	transactions, err := repo.AdminTransactionsFindAllByAdminID(context.TODO(), adminID)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, transactions, util.NewStatus(http.StatusOK, ""))
}

func AdminTransactionsGet(response http.ResponseWriter, request *http.Request) {
	transactions, err := repo.AdminTransactionsFindAll(context.TODO())

	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	// adapters := make([]*adapterPayment, 0, len(transactions))
	// for _, it := range transactions {
	// 	adapters = append(adapters, paymentToAdapter(it))
	// }

	// util.Response(response, adapters, util.NewStatus(http.StatusOK, ""))
	util.Response(response, transactions, util.NewStatus(http.StatusOK, ""))
}
