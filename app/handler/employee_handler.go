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
)

func EmployeesGet(response http.ResponseWriter, request *http.Request) {
	employees, err := repo.EmployeesFindAll(context.TODO())
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusInternalServerError, err.Error()))
		return
	}

	util.Response(response, employees, util.NewStatus(http.StatusOK, ""))
}

func EmployeeGet(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)

	id, _ := params["id"]
	employee, err := repo.EmployeesFindOneByID(context.TODO(), id)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}
	util.Response(response, employee, util.NewStatus(http.StatusOK, ""))
}

func EmployeeGetByEmployeeID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	employeeID, _ := params["employee_id"]

	employee, err := repo.EmployeesFindByEmployeeID(context.TODO(), employeeID)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, err.Error()))
		return
	}
	util.Response(response, employee, util.NewStatus(http.StatusOK, ""))
}

func EmployeePost(response http.ResponseWriter, request *http.Request) {
	var employee model.Employee
	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	if err := employee.Validate(); err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	// Check if employee with employee_id already exists, if it does dont add that employee
	if repo.EmployeesIsAlreadyExistsWithEmployeeID(context.TODO(), employee.EmployeeID) {
		httpError := util.NewStatus(http.StatusForbidden, "employee already exists")
		util.Response(response, struct{}{}, httpError)
		return
	}

	e, err := repo.EmployeesInsertOne(context.TODO(), &employee)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusInternalServerError, err.Error()))
		return
	}
	util.Response(response, e, util.NewStatus(http.StatusOK, ""))
}

func EmployeePutByEmployeeID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	employee_id, _ := params["employee_id"]

	var employee model.Employee
	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		util.Response(response, struct{}{}, util.NewStatus(http.StatusBadRequest, err.Error()))
		return
	}

	if employee_id != employee.EmployeeID {
		httpError := util.NewStatus(http.StatusBadRequest, "employee_id update and employee employee_id mismatch")
		util.Response(response, struct{}{}, httpError)
		return
	}

	if err := employee.Validate(); err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}

	// Check if employee with employee_id already exists, if it does update
	if !repo.EmployeesIsAlreadyExistsWithEmployeeID(context.TODO(), employee.EmployeeID) {
		httpError := util.NewStatus(http.StatusForbidden, "employee doesnt exist to update")
		util.Response(response, struct{}{}, httpError)
		return
	}

	e, err := repo.EmployeesUpdateOneByEmployeeID(context.TODO(), employee_id, &employee)
	if err != nil {
		httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, e, util.NewStatus(http.StatusOK, ""))
}

func EmployeeDelete(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	internalId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		httpError := util.NewStatus(http.StatusBadRequest, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	employee := &model.Employee{ID: internalId}
	e, err := repo.EmployeesDeleteOne(context.TODO(), employee)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, e, util.NewStatus(http.StatusOK, ""))
}

func EmployeeDeleteByEmployeeID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	employee_id, _ := params["employee_id"]

	employee, err := repo.EmployeesDeleteByEmployeeID(context.TODO(), employee_id)
	if err != nil {
		httpError := util.NewStatus(http.StatusNotFound, err.Error())
		util.Response(response, struct{}{}, httpError)
		return
	}
	util.Response(response, employee, util.NewStatus(http.StatusOK, ""))
}
