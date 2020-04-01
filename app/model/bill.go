package model

import (
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"errors"
)

// MealBill is bill for each mealtype
type MealBill struct {
	Quantity int     `json:"quantity,required" bson:"quantity,required"`
	Total    float64 `json:"total,required" bson:"total,required"`
}

// Validate validates mealbill
func (m *MealBill) Validate() error {
	if m.Quantity < 0 {
		return errors.New("quantity is negative")
	}

	if m.Total < 0 {
		return errors.New("total is negative")
	}

	return nil
}

// Bill is overall bill
type Bill struct {
	User      User     `json:"user,required" bson:"user,required"`
	Breakfast MealBill `json:"breakfast,required" bson:"breakfast,required"`
	Lunch     MealBill `json:"lunch,required" bson:"lunch,required"`
	Snack     MealBill `json:"snack,required" bson:"snack,required"`
	TotalDue  float64  `json:"total,required" bson:"total,required"`
}

// Validate receipt fields
// This function validates receipt data
// and return error is any
// all errors are related to the fields
func (b *Bill) Validate() error {
	// validating user field with retuired, min length 1, max length 25 and regex check

	if e := b.User.Validate(); e != nil {
		return e
	}

	if e := b.Breakfast.Validate(); e != nil {
		return e
	}

	if e := b.Lunch.Validate(); e != nil {
		return e
	}

	if e := b.Snack.Validate(); e != nil {
		return e
	}

	return nil
}
