package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
)

// MealType model
type MealType struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type      string             `json:"type,omitempty" bson:"type,omitempty"`
	Cost      float64            `json:"cost,required" bson:"cost,required"`
	StartTime primitive.DateTime `json:"start_time,required" bson:"start_time,required"`
	EndTime   primitive.DateTime `json:"end_time,required" bson:"end_time,required"`
}

// Validate fields
// This function validates meal data
// and return error is any
// all errors are related to the fields
func (m *MealType) Validate() error {
	// validating type field with retuired, min length 1, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(m.Type, true, 1, 25, "[a-zA-Z0-9]+", "type"); e != nil {
		return e
	}

	// validating cost field with required, min length 0, max length 0 and regex check
	if m.Cost < 0 {
		return errors.New("cost is negative")
	}

	// validating email field with required, min length 5, max length 25 and regex check
	if m.EndTime < m.StartTime {
		return errors.New("start_time is greater than end_time")
	}

	return nil
}
