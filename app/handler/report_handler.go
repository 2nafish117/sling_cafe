package handler

import (
	"context"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	// "sling_cafe/app/model"

	"sling_cafe/app/repo"
	"sling_cafe/util"
	// "time"
)

// per employee meal entry report with number of each meals
// aggregatte the meal entries and get total

// MealEntryReportsGet get report of all employees
func MealEntryReportsGet(response http.ResponseWriter, request *http.Request) {
	reports, err := repo.MealEntriesAggregateReports(context.TODO())
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, reports, util.NewStatus(http.StatusOK, ""))
}

func MealEntryReportGetByEmployeeID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	employeeID, _ := params["employee_id"]
	report, err := repo.MealEntriesAggregateReportByEmployeeID(context.TODO(), employeeID)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, report, util.NewStatus(http.StatusOK, ""))
}

// per employee transactions report with balance to pay
// aggregation of employee monthly posting - employee payments transactions
func EmployeeTransactionReportsGet(response http.ResponseWriter, request *http.Request) {
	reports, err := repo.EmployeeTransactionsAggregateReports(context.TODO())
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, reports, util.NewStatus(http.StatusOK, ""))
}

// employee transaction report for one employee based on employee_id
func EmployeeTransactionReportGetByEmployeeID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	employeeID, _ := params["employee_id"]
	report, err := repo.EmployeeTransactionsAggregateReportByEmployeeID(context.TODO(), employeeID)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, report, util.NewStatus(http.StatusOK, ""))
}

// admins transactions report with cash in hand
// aggregation of employee payments - admin deposits

func AdminTransactionReportsGet(response http.ResponseWriter, request *http.Request) {
	reports, err := repo.AdminTransactionsAggregateReports(context.TODO())
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, reports, util.NewStatus(http.StatusOK, ""))
}

func AdminTransactionReportGetByAdminID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	adminID, _ := params["admin_id"]
	report, err := repo.AdminTransactionsAggregateReportByAdminID(context.TODO(), adminID)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, report, util.NewStatus(http.StatusOK, ""))
}

// per caterer transactions report  with balance to pay
//

// monthly summary report -
//
