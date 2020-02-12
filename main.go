package main

import (
	"context"
	// "encoding/json"
	"fmt"
	"sling_cafe/config"
	"sling_cafe/db"

	"net/http"
	// "time"

	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/app/handler"
)

func main() {
	fmt.Println("Starting the application...")

	cfg := config.GetInstance()
	fmt.Println(cfg)
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
	fmt.Println(base)
	router.HandleFunc(base+"/users", handler.UserPost).Methods("POST")
	router.HandleFunc(base+"/users/{id:[0-9a-f]{24}}", handler.UserGet).Methods("GET") // order of these registrations matter !!!
	router.HandleFunc(base+"/users/{empid:[a-zA-Z0-9]+}", handler.UserGetByEmpid).Methods("GET")
	// @TODO: Add filters, pagination
	router.HandleFunc(base+"/users", handler.UsersGet).Methods("GET")
	router.HandleFunc(base+"/users/{id:[0-9a-f]{24}}", handler.UserDelete).Methods("DELETE") // order of these registrations matter !!!
	router.HandleFunc(base+"/users/{empid:[a-zA-Z0-9]+}", handler.UserDeleteByEmpid).Methods("DELETE")

	router.HandleFunc(base+"/meals", handler.MealPost).Methods("POST")
	// @TODO: Add filters, pagination
	router.HandleFunc(base+"/meals", handler.MealsGet).Methods("GET")
	router.HandleFunc(base+"/meals/user/{empid:[a-zA-Z0-9]+}", handler.MealsGetByEmpid).Methods("GET")

	router.HandleFunc(base+"/mealtypes", handler.MealTypePost).Methods("POST")
	// @TODO: Add filters, pagination
	router.HandleFunc(base+"/mealtypes", handler.MealTypesGet).Methods("GET")

	// router.HandleFunc(base+"/receipts/{id:[0-9a-f]{24}", handler.ReceiptGet).Methods("GET")
	// router.HandleFunc(base+"/receipts/", handler.ReceiptsGet).Methods("GET")
	// @TODO listen and serve arguments fix?
	http.ListenAndServe(cfg.ApiAddr, router)
}
