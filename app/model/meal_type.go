package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "sling_cafe/util"
	"time"
)

// MealType model
type MealType struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MealID       MealID             `json:"meal_id,required" bson:"meal_id,required"`
	Cost         int64              `json:"cost,required,string" bson:"cost,required"`
	CompanyCost  int64              `json:"company_cost,required,string" bson:"company_cost,required"`
	EmployeeCost int64              `json:"employee_cost,required,string" bson:"employee_cost,required"`
	FromTime     time.Time          `json:"from_time,required" bson:"from_time,required"`
	ToTime       time.Time          `json:"to_time,required" bson:"to_time,required"`
	CatererID    string             `json:"caterer_id,required" bson:"caterer_id,required"` // who served that meal?
	Inactive     bool               `json:"inactive,required" bson:"inactive,required"`     // is mealtype inactive ?
}

// Validate fields
// This function validates meal data
// and return error is any
// all errors are related to the fields
func (m *MealType) Validate() error {
	// validating type field with retuired, min length 1, max length 25 and regex check
	// if e := util.ValidateRequireAndLengthAndRegex(m.MealID, true, 1, 25, "[a-zA-Z0-9]+", "type"); e != nil {
	// 	return e
	// }

	// validating cost field with required, min length 0, max length 0 and regex check
	if m.Cost < 0 {
		return errors.New("cost is negative")
	}

	// validating cost field with required, min length 0, max length 0 and regex check
	if m.CompanyCost < 0 {
		return errors.New("company_cost is negative")
	}

	// validating cost field with required, min length 0, max length 0 and regex check
	if m.EmployeeCost < 0 {
		return errors.New("employee_cost is negative")
	}

	if m.EmployeeCost+m.CompanyCost != m.Cost {
		return errors.New("employee cost + company cost != total cost")
	}

	if m.ToTime.Unix() < m.FromTime.Unix() {
		return errors.New("from_time should be before to_time")
	}

	return nil
}
