package main

import (
	"context"
	"sling_cafe/app/repo"
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

	base := "/" + cfg.ApiName + "/" + cfg.ApiVersion
	// @TODO: add a search endpoint, searching fullnames dont work !!
	router.HandleFunc(base+"/search/{resource:[-_a-zA-Z]+}", handler.SearchHandler).Methods("GET")

	router.HandleFunc(base+"/employees", handler.EmployeePost).Methods("POST")
	router.HandleFunc(base+"/employees/{id:[0-9a-f]{24}}", handler.EmployeeGet).Methods("GET") // order of these registrations matter !!!
	router.HandleFunc(base+"/employees/{employee_id:[a-zA-Z0-9]+}", handler.EmployeeGetByEmployeeID).Methods("GET")
	router.HandleFunc(base+"/employees", handler.EmployeesGet).Methods("GET")
	router.HandleFunc(base+"/employees/{employee_id:[a-zA-Z0-9]+}", handler.EmployeePutByEmployeeID).Methods("PUT")
	router.HandleFunc(base+"/employees/{id:[0-9a-f]{24}}", handler.EmployeeDelete).Methods("DELETE") // order of these registrations matter !!!
	router.HandleFunc(base+"/employees/{employee_id:[a-zA-Z0-9]+}", handler.EmployeeDeleteByEmployeeID).Methods("DELETE")

	router.HandleFunc(base+"/admins", handler.AdminPost).Methods("POST")
	router.HandleFunc(base+"/admins/{id:[0-9a-f]{24}}", handler.AdminGet).Methods("GET") // order of these registrations matter !!!
	router.HandleFunc(base+"/admins/{admin_id:[a-zA-Z0-9]+}", handler.AdminGetByAdminID).Methods("GET")
	router.HandleFunc(base+"/admins", handler.AdminsGet).Methods("GET")
	router.HandleFunc(base+"/admins/{admin_id:[a-zA-Z0-9]+}", handler.AdminPutByAdminID).Methods("PUT")
	router.HandleFunc(base+"/admins/{id:[0-9a-f]{24}}", handler.AdminDelete).Methods("DELETE") // order of these registrations matter !!!
	router.HandleFunc(base+"/admins/{admin_id:[a-zA-Z0-9]+}", handler.AdminDeleteByAdminID).Methods("DELETE")

	router.HandleFunc(base+"/caterers", handler.CatererPost).Methods("POST")
	router.HandleFunc(base+"/caterers/{id:[0-9a-f]{24}}", handler.CatererGet).Methods("GET") // order of these registrations matter !!!
	router.HandleFunc(base+"/caterers/{caterer_id:[a-zA-Z0-9]+}", handler.CatererGetByCatererID).Methods("GET")
	router.HandleFunc(base+"/caterers", handler.CaterersGet).Methods("GET")
	router.HandleFunc(base+"/caterers/{caterer_id:[a-zA-Z0-9]+}", handler.CatererPutByCatererID).Methods("PUT")
	router.HandleFunc(base+"/caterers/{id:[0-9a-f]{24}}", handler.CatererDelete).Methods("DELETE") // order of these registrations matter !!!
	router.HandleFunc(base+"/caterers/{caterer_id:[a-zA-Z0-9]+}", handler.CatererDeleteByCatererID).Methods("DELETE")

	router.HandleFunc(base+"/meal_entries", handler.MealEntryPost).Methods("POST")
	router.HandleFunc(base+"/meal_entries", handler.MealEntriesGet).Methods("GET")
	router.HandleFunc(base+"/meal_entries/employee/{employee_id:[a-zA-Z0-9]+}", handler.MealEntriesGetByEmployeeID).Methods("GET")

	router.HandleFunc(base+"/employee_transactions", handler.EmployeeTransactionPost).Methods("POST")
	router.HandleFunc(base+"/employee_transactions", handler.EmployeeTransactionsGet).Methods("GET")
	router.HandleFunc(base+"/employee_transactions/employee/{employee_id:[a-zA-Z0-9]+}", handler.EmployeeTransactionsGetByEmployeeID).Methods("GET")

	router.HandleFunc(base+"/admin_transactions", handler.AdminTransactionPost).Methods("POST")
	router.HandleFunc(base+"/admin_transactions", handler.AdminTransactionsGet).Methods("GET")
	router.HandleFunc(base+"/admin_transactions/admin/{admin_id:[a-zA-Z0-9]+}", handler.AdminTransactionsGetByAdminID).Methods("GET")

	router.HandleFunc(base+"/meal_types", handler.MealTypePost).Methods("POST")
	router.HandleFunc(base+"/meal_types", handler.MealTypesGet).Methods("GET")
	router.HandleFunc(base+"/meal_types/{meal_id}", handler.MealTypeGet).Methods("GET")
	router.HandleFunc(base+"/meal_types/{meal_id}", handler.MealTypePutByMealID).Methods("PUT")
	router.HandleFunc(base+"/meal_types/{meal_id}", handler.MealTypeDeleteByMealID).Methods("DELETE")

	// @TODO:
	router.HandleFunc(base+"/reports/meal_entries", handler.MealEntryReportsGet).Methods("GET")
	router.HandleFunc(base+"/reports/meal_entries/employee/{employee_id:[a-zA-Z0-9]+}", handler.MealEntryReportGetByEmployeeID).Methods("GET")
	router.HandleFunc(base+"/reports/employee_transactions", handler.EmployeeTransactionReportsGet).Methods("GET")
	router.HandleFunc(base+"/reports/employee_transactions/employee/{employee_id:[a-zA-Z0-9]+}", handler.EmployeeTransactionReportGetByEmployeeID).Methods("GET")
	router.HandleFunc(base+"/reports/admin_transactions", handler.AdminTransactionReportsGet).Methods("GET")
	router.HandleFunc(base+"/reports/admin_transactions/admin/{admin_id:[a-zA-Z0-9]+}", handler.AdminTransactionReportGetByAdminID).Methods("GET")

	// router.HandleFunc(base+"/reports/caterer_transactions", handler.CatererTransactionReportsGet).Methods("GET")
	// router.HandleFunc(base+"/reports/caterer_transactions/caterer/{caterer_id:[a-zA-Z0-9]+}", handler.CatererTransactionReportGetByCatererID).Methods("GET")

	// go repo.MonthEndEvent()
	// go repo.DayEndEvent()
	// go repo.HourEndEvent()
	go repo.MinuteEndEvent()
	// go repo.SecondEndEvent()

	handler := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(router)

	Log.Info("Application is running at: ", cfg.ApiAddr+base)
	http.ListenAndServe(cfg.ApiAddr, handler)
}
