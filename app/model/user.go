package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
)

// User model
type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UId   string             `json:"uid,required" bson:"uid,required"`
	Fname string             `json:"fname,required" bson:"fname,required"`
	Lname string             `json:"lname,omitempty" bson:"lname,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
}

// Validate user fields
// This function validates user data
// and return error is any
// all errors are related to the fields
func (u *User) Validate() error {

	// validating uid field with retuired, min length 1, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.UId, true, 1, 25, "[a-zA-Z0-9]+", "uid"); e != nil {
		return e
	}

	// validating firstname field with retuired, min length 3, max length 25 and no regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.Fname, true, 3, 25, "", "fname"); e != nil {
		return e
	}

	// validating lastname field with retuired, min length 0, max length 25 and no regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.Lname, false, 3, 25, "", "lname"); e != nil {
		return e
	}

	// validating email field with required, min length 5, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.Email, true, 5, 25, `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, "email"); e != nil {
		return e
	}

	return nil
}
