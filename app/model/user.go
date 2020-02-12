package model

import (
	// "context"
	// "encoding/json"
	// "fmt"

	// "github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
	// "mongo_test/db"
	// "net/http"
	// "time"
)

// User model
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmpId     string             `json:"empid,required" bson:"empid,required"`
	Firstname string             `json:"firstname,required" bson:"firstname,required"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
}

type UserResponse struct {
	Empid string `json:"empid,required" bson:"empid,required"`
}

// Validate user fields
// This function validates user data
// and return error is any
// all errors are related to the fields
func (u *User) Validate() error {

	// @TODO: add regex checks!!
	// validating firstname field with retuired, min length 3, max length 25 and no regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.Firstname, true, 3, 25, "", "firstname"); e != nil {
		return e
	}

	// validating lastname field with retuired, min length 0, max length 25 and no regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.Lastname, false, 3, 25, "", "lastname"); e != nil {
		return e
	}

	// validating email field with required, min length 5, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.Email, true, 5, 25, `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, "email"); e != nil {
		return e
	}

	return nil
}
