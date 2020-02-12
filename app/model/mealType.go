package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
   "mealtypeid": "snacksid",
   "menutypeid": "bondabajjiid",
   "starttime": "start_time_stamp_for_snacks",
   "endtime": "end_time_for_snacks",
   "costtocompany": 234,
   "costtoemployee": 100

*/

// @ TODO: see how to handle storing of starttime and endtime !!!!!!!
// MealType holds data about each possible meal type, eg: breakfast, lunch, snacks, dinner ...
type MealType struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MealTypeId    string             `json:"mealtypeid,required" bson:"mealtypeid,required"`
	StartTime     string             `json:"starttime,required" bson:"starttime,required"`
	EndTime       string             `json:"endtime,required" bson:"endtime,required"`
	CostToCompany float32            `json:"costtocompany,required" bson:"costtocompany,required"`
	CostToUser    float32            `json:"costtouser,required" bson:"costtouser,required"`
}

// type MealType struct {
// 	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	MealTypeId    string             `json:"mealtypeid,required" bson:"mealtypeid,required"`
// 	StartTime     primitive.DateTime `json:"starttime,required" bson:"starttime,required"`
// 	EndTime       primitive.DateTime `json:"endtime,required" bson:"endtime,required"`
// 	CostToCompany float32            `json:"costtocompany,required" bson:"costtocompany,required"`
// 	CostToUser    float32            `json:"costtouser,required" bson:"costtouser,required"`
// }

// Validate fields
// This function validates mealtype data
// and return error is any
// all errors are related to the fields
func (m *MealType) Validate() error {

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
