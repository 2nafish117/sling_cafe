package handler

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	// "fmt"
	"github.com/gorilla/mux"
	"sling_cafe/app/repo"
	"sling_cafe/util"
)

// http://localhost:12345/slingcafe/v1/search/resource-type?querystring
// db.stores.find( { $text: { $search: "java coffee shop" } } )

// SearchHandler searches based on text, names in resoureces
func SearchHandler(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	queries := request.URL.Query()

	resource, _ := params["resource"]
	q := queries.Get("q")

	dbQuery := bson.M{"$text": bson.M{"$search": q}}

	switch resource {
	case "employees":
		employees, err := repo.EmployeesFind(context.TODO(), dbQuery)
		if err != nil {
			httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
			util.Response(response, struct{}{}, httpError)
			return
		}
		util.Response(response, employees, util.NewStatus(http.StatusOK, ""))
		return
	case "admins":
		admins, err := repo.AdminsFind(context.TODO(), dbQuery)
		if err != nil {
			httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
			util.Response(response, struct{}{}, httpError)
			return
		}
		util.Response(response, admins, util.NewStatus(http.StatusOK, ""))
		return
	case "caterers":
		caterers, err := repo.CaterersFind(context.TODO(), dbQuery)
		if err != nil {
			httpError := util.NewStatus(http.StatusInternalServerError, err.Error())
			util.Response(response, struct{}{}, httpError)
			return
		}
		util.Response(response, caterers, util.NewStatus(http.StatusOK, ""))
		return
	default:
		util.Response(response, struct{}{}, util.NewStatus(http.StatusNotFound, "resource type not searchable"))
		return
	}
}
