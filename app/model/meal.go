package model

import (
	// "context"
	// "encoding/json"
	// "fmt"

	// "github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "sling_cafe/util"
	// "mongo_test/db"
	// "net/http"
	// "time"
)

/*
{
    "mealtypeid": "lunchid",
    "menuitemid" : "southindianid",
    "timestamp": "sometime",
    "employeeid": "dsfg9873"
}

*/
// Meal model
type Meal struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MealTypeId string             `json:"mealtypeid,required" bson:"mealtypeid,required"` // reference to a meal type
	Timestamp  primitive.DateTime `json:"timestamp,required" bson:"timestamp,required"`   // time is generated server side using time.Now()
	EmpId      string             `json:"empid,omitempty" bson:"empid,omitempty"`         // reference to an employee
}

// Validate fields
// This function validates meal data
// and return error is any
// all errors are related to the fields
func (m *Meal) Validate() error {

	// @TODO: add regex checks!!
	// validating firstname field with retuired, min length 3, max length 25 and no regex check
	// if e := util.ValidateRequireAndLengthAndRegex(m.Firstname, true, 3, 25, "", "firstname"); e != nil {
	// 	return e
	// }

	// // validating lastname field with retuired, min length 0, max length 25 and no regex check
	// if e := util.ValidateRequireAndLengthAndRegex(m.Lastname, false, 3, 25, "", "lastname"); e != nil {
	// 	return e
	// }

	// // validating email field with required, min length 5, max length 25 and regex check
	// if e := util.ValidateRequireAndLengthAndRegex(m.Email, true, 5, 25, `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, "email"); e != nil {
	// 	return e
	// }

	return nil
}

// var empid = ObjectId("5e3cda4d656f6acb7c10f61f")
// db.meals.insert({mealtype : "lunch", timestamp: new Date(), "empid" : empid})
