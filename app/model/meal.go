package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
)

// Meal model
type Meal struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UId      string             `json:"uid,omitempty" bson:"uid,omitempty"`         // reference to an employee
	DateTime primitive.DateTime `json:"datetime,required" bson:"datetime,required"` // time is generated server side using time.Now()
	Type     string             `json:"type,required" bson:"type,required"`         // reference to a meal type
	Cost     float64            `json:"cost,required" bson:"cost,required"`
}

// Validate fields
// This function validates meal data
// and return error is any
// all errors are related to the fields
func (m *Meal) Validate() error {
	// validating uid field with retuired, min length 1, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(m.UId, true, 1, 25, "[a-zA-Z0-9]+", "uid"); e != nil {
		return e
	}

	// validating email field with required, min length 5, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(m.Type, true, 5, 25, `[a-z]+`, "type"); e != nil {
		return e
	}

	// validating cost field with required, min length 0, max length 0 and regex check
	if m.Cost < 0 {
		return errors.New("cost is negative")
	}

	return nil
}
