package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
)

// Meal model
type Meal struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmpId     string             `json:"empid,omitempty" bson:"empid,omitempty"` // reference to an employee
	Timestamp primitive.DateTime `json:"time,required" bson:"time,required"`     // time is generated server side using time.Now()
	Cost      float64            `json:"cost,required" bson:"cost,required"`
	MealType  string             `json:"mealtype,required" bson:"mealtype,required"` // reference to a meal type
}

// Validate fields
// This function validates meal data
// and return error is any
// all errors are related to the fields
func (m *Meal) Validate() error {
	// validating empid field with retuired, min length 1, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(m.EmpId, true, 1, 25, "[a-zA-Z0-9]+", "empid"); e != nil {
		return e
	}

	// validating email field with required, min length 5, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(m.MealType, true, 5, 25, `[a-z]+`, "mealtype"); e != nil {
		return e
	}

	// validating cost field with required, min length 0, max length 0 and regex check
	if m.Cost < 0 {
		return errors.New("cost is negative")
	}

	return nil
}
