package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Receipt struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmpId  string             `json:"empid,required" bson:"empid,required"`
	AmtDue float32            `json:"amtdue,required" bson:"amtdue,required"`
}

// Validate receipt fields
// This function validates receipt data
// and return error is any
// all errors are related to the fields
func (r *Receipt) Validate() error {

	// @TODO: add regex checks!!
	// validating firstname field with retuired, min length 3, max length 25 and no regex check
	// if e := util.ValidateRequireAndLengthAndRegex(u.Firstname, true, 3, 25, "", "firstname"); e != nil {
	// 	return e
	// }

	// // validating lastname field with retuired, min length 0, max length 25 and no regex check
	// if e := util.ValidateRequireAndLengthAndRegex(u.Lastname, false, 3, 25, "", "lastname"); e != nil {
	// 	return e
	// }

	// // validating email field with required, min length 5, max length 25 and regex check
	// if e := util.ValidateRequireAndLengthAndRegex(u.Email, true, 5, 25, `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, "email"); e != nil {
	// 	return e
	// }

	return nil
}
