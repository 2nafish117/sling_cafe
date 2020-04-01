package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
)

type Receipt struct {
	User     User               `json:"user,omitempty" bson:"user,omitempty"`
	Mode     string             `json:"mode,required" bson:"mode,required"`
	Amount   float64            `json:"amount,required" bson:"amount,required"`
	DateTime primitive.DateTime `json:"datetime,required" bson:"datetime,required"`
}

// Validate receipt fields
// This function validates receipt data
// and return error is any
// all errors are related to the fields
func (r *Receipt) Validate() error {
	if e := r.User.Validate(); e != nil {
		return e
	}

	if e := util.ValidateRequireAndLengthAndRegex(r.Mode, true, 1, 25, "[a-zA-Z ]+", "mode"); e != nil {
		return e
	}

	if r.Amount < 0 {
		return errors.New("amount is negative")
	}

	return nil
}
