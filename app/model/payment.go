package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
)

// Payment model
type Payment struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UId      string             `json:"uid,omitempty" bson:"uid,omitempty"`         // reference to an employee
	DateTime primitive.DateTime `json:"datetime,required" bson:"datetime,required"` // time is generated server side using time.Now()
	Mode     string             `json:"mode,required" bson:"mode,required"`         // mode of payment
	Amount   float64            `json:"amount,required" bson:"amount,required"`
}

// Validate fields
// This function validates meal data
// and return error is any
// all errors are related to the fields
func (p *Payment) Validate() error {
	// validating uid field with retuired, min length 1, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(p.UId, true, 1, 25, "[a-zA-Z0-9]+", "uid"); e != nil {
		return e
	}

	if e := util.ValidateRequireAndLengthAndRegex(p.Mode, true, 1, 25, "[a-zA-Z ]+", "mode"); e != nil {
		return e
	}

	// validating cost field with required, min length 0, max length 0 and regex check
	if p.Amount < 0 {
		return errors.New("amount is negative")
	}

	return nil
}
