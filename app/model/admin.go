package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
	"time"
)

// Admin model
type Admin struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AdminID   string             `json:"admin_id,required" bson:"admin_id,required"`
	Name      string             `json:"name,required" bson:"name,required"`
	Contact   string             `json:"contact,required" bson:"contact,required"`
	Email     string             `json:"email,required" bson:"email,required"`
	JoinDate  time.Time          `json:"join_date,omitempty" bson:"join_date,omitempty"`
	LeaveDate time.Time          `json:"leave_date,omitempty" bson:"leave_date,omitempty"`
	UserName  string             `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
}

// Validate user fields
// This function validates user data
// and return error is any
// all errors are related to the fields
func (u *Admin) Validate() error {

	// validating uid field with retuired, min length 1, max length 200 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.AdminID, true, 1, 200, "[a-zA-Z0-9]+", "admin_id"); e != nil {
		return e
	}

	// validating firstname field with retuired, min length 3, max length 200 and no regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.Name, true, 2, 200, "", "name"); e != nil {
		return e
	}

	// validating contact field with required, min length 5, max length 200 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.Contact, true, 5, 200, `^[0-9]{10}$`, "contact"); e != nil {
		return e
	}

	// validating email field with required, min length 5, max length 200 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.Email, true, 5, 200, `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, "email"); e != nil {
		return e
	}

	// validating firstname field with retuired, min length 3, max length 200 and no regex check
	if e := util.ValidateRequireAndLengthAndRegex(u.UserName, false, 2, 200, "", "user_name"); e != nil {
		return e
	}

	return nil
}
