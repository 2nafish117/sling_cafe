package main

import (
	"context"
	// "fmt"
	"sling_cafe/config"
	"sling_cafe/db"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"sling_cafe/app/handler"
	. "sling_cafe/log"
)

func main() {
	Log.Info("Starting...")

	cfg := config.GetInstance()
	db.Connect()
	db := db.GetInstance()
	defer db.Disconnect(context.TODO())

	router := mux.NewRouter()

	/*
		POST
		GET
		PUT
		DELETE
	*/

	base := "/" + cfg.ApiName + "/" + cfg.ApiVersion
	router.HandleFunc(base+"/users", handler.UserPost).Methods("POST")
	router.HandleFunc(base+"/users/{id:[0-9a-f]{24}}", handler.UserGet).Methods("GET") // order of these registrations matter !!!
	router.HandleFunc(base+"/users/{empid:[a-zA-Z0-9]+}", handler.UserGetByEmpid).Methods("GET")
	// @TODO: Add filters, pagination
	router.HandleFunc(base+"/users", handler.UsersGet).Methods("GET")
	router.HandleFunc(base+"/users/{empid:[a-zA-Z0-9]+}", handler.UserPutByEmpid).Methods("PUT")
	router.HandleFunc(base+"/users/{id:[0-9a-f]{24}}", handler.UserDelete).Methods("DELETE") // order of these registrations matter !!!
	router.HandleFunc(base+"/users/{empid:[a-zA-Z0-9]+}", handler.UserDeleteByEmpid).Methods("DELETE")

	router.HandleFunc(base+"/meals", handler.MealPost).Methods("POST")
	// @TODO: Add filters, pagination
	router.HandleFunc(base+"/meals", handler.MealsGet).Methods("GET")
	router.HandleFunc(base+"/meals/user/{empid:[a-zA-Z0-9]+}", handler.MealsGetByEmpid).Methods("GET")

	// @TODO: Add filters, pagination
	router.HandleFunc(base+"/receipts", handler.ReceiptsGet).
		Queries("start", "{start}", "end", "{end}").Methods("GET")
	router.HandleFunc(base+"/receipts/user/{empid:[a-zA-Z0-9]+}", handler.ReceiptGetByEmpid).
		Queries("start", "{start}", "end", "{end}").Methods("GET")
	router.HandleFunc(base+"/receipts/user/{id:[0-9a-f]{24}", handler.ReceiptGetById).
		Queries("start", "{start}", "end", "{end}").Methods("GET")

	router.HandleFunc(base+"/vendors", handler.VendorPost).Methods("POST")
	router.HandleFunc(base+"/vendors/{id:[0-9a-f]{24}}", handler.VendorGet).Methods("GET") // order of these registrations matter !!!
	router.HandleFunc(base+"/vendors/{vid:[a-zA-Z0-9]+}", handler.VendorGetByVid).Methods("GET")
	// @TODO: Add filters, pagination
	router.HandleFunc(base+"/vendors", handler.VendorsGet).Methods("GET")
	router.HandleFunc(base+"/vendors/{vid:[a-zA-Z0-9]+}", handler.VendorPutByVid).Methods("PUT")
	router.HandleFunc(base+"/vendors/{id:[0-9a-f]{24}}", handler.VendorDelete).Methods("DELETE") // order of these registrations matter !!!
	router.HandleFunc(base+"/vendors/{vid:[a-zA-Z0-9]+}", handler.VendorDeleteByVid).Methods("DELETE")

	handler := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(router)

	Log.Info("Application is running at: ", cfg.ApiAddr+base)
	http.ListenAndServe(cfg.ApiAddr, handler)
}
