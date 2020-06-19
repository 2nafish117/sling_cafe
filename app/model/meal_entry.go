package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
	"time"
)

// MealEntry model
type MealEntry struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DateTime     time.Time          `json:"date_time,omitempty" bson:"date_time,omitempty"` // time is generated server side using time.Now()
	EmployeeID   string             `json:"employee_id,required" bson:"employee_id,required"`
	MealID       MealID             `json:"meal_id,required" bson:"meal_id,required"`
	CatererID    string             `json:"caterer_id,omitempty" bson:"caterer_id,omitempty"`              // who is the currently working caterer, deduced server side
	Cost         int64              `json:"cost,omitempty,string" bson:"cost,omitempty"`                   // cost of this meal, deduced server side
	EmployeeCost int64              `json:"employee_cost,omitempty,string" bson:"employee_cost,omitempty"` // cost of this meal, deduced server side
	CompanyCost  int64              `json:"company_cost,omitempty,string" bson:"company_cost,omitempty"`   // cost of this meal, deduced server side
}

// Validate fields
// This function validates meal data
// and return error is any
// all errors are related to the fields
func (m *MealEntry) Validate() error {
	// validating uid field with retuired, min length 1, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(m.EmployeeID, true, 1, 25, "[a-zA-Z0-9]+", "employee_id"); e != nil {
		return e
	}

	// validating email field with required, min length 5, max length 25 and regex check
	// if e := util.ValidateRequireAndLengthAndRegex(m.MealID, false, 5, 25, `[a-z]+`, "meal_id"); e != nil {
	// 	return e
	// }

	// validating cost field with required, min length 0, max length 0 and regex check
	if m.Cost < 0 {
		return errors.New("cost is negative")
	}

	return nil
}

type MealEntryReport struct {
	EmployeeID   string `json:"employee_id,required" bson:"employee_id,required"`
	Name         string `json:"name,required" bson:"name,required"`
	BreakfastQty int    `json:"breakfast_quantity,required" bson:"breakfast_quantity,required"`
	LunchQty     int    `json:"lunch_quantity,required" bson:"lunch_quantity,required"`
	SnackQty     int    `json:"snack_quantity,required" bson:"snack_quantity,required"`
}

func (b *MealEntryReport) Validate() error {

	if b.BreakfastQty < 0 {
		return errors.New("breakfast_quantity is negative")
	}

	if b.LunchQty < 0 {
		return errors.New("lunch_quantity is negative")
	}

	if b.SnackQty < 0 {
		return errors.New("snack_quantity is negative")
	}

	return nil
}
