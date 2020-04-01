package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
)

// Vendor model
type Vendor struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VId       string             `json:"vid,required" bson:"vid,required"`
	Firstname string             `json:"fname,required" bson:"fname,required"`
	Lastname  string             `json:"lname,omitempty" bson:"lname,omitempty"`
}

// Validate user fields
// This function validates user data
// and return error is any
// all errors are related to the fields
func (v *Vendor) Validate() error {

	// validating uid field with retuired, min length 1, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(v.VId, true, 1, 25, "[a-zA-Z0-9]+", "vid"); e != nil {
		return e
	}

	// validating firstname field with retuired, min length 3, max length 25 and no regex check
	if e := util.ValidateRequireAndLengthAndRegex(v.Firstname, true, 3, 25, "", "fname"); e != nil {
		return e
	}

	// validating lastname field with retuired, min length 0, max length 25 and no regex check
	if e := util.ValidateRequireAndLengthAndRegex(v.Lastname, false, 3, 25, "", "lname"); e != nil {
		return e
	}

	return nil
}
