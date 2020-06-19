package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
	"time"
)

// @TODO: to calculate the owed amount of an employee,
// every month end a posting is created in this table, which is total cost eaten by then employee in that month.
// everytime an employee pays an entry is created in this table.
// sum of all postings minus sum of all payments from this table is the amount due.
// admin_id for postings is blank

// EmployeeTransaction model
type EmployeeTransaction struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TransactionID   string             `json:"transaction_id,required" bson:"transaction_id,required"`
	DateTime        time.Time          `json:"date_time,required" bson:"date_time,required"`               // time is generated server side using time.Now()
	TransactionType TransactionType    `json:"transaction_type,required" bson:"transaction_type,required"` // can be of 2 types, monthly posting or employee payment
	EmployeeID      string             `json:"employee_id,required" bson:"employee_id,required"`
	AdminID         string             `json:"admin_id,omitempty" bson:"admin_id,omitempty"`
	Mode            Mode               `json:"mode,omitempty" bson:"mode,omitempty"`
	Amount          int64              `json:"amount,required,string" bson:"amount,required"`
	Remarks         string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
}

// Validate fields
// This function validates meal data
// and return error is any
// all errors are related to the fields
func (p *EmployeeTransaction) Validate() error {
	// validating uid field with retuired, min length 1, max length 25 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(p.EmployeeID, true, 1, 25, "[a-zA-Z0-9]+", "employee_id"); e != nil {
		return e
	}

	if e := util.ValidateRequireAndLengthAndRegex(p.AdminID, true, 1, 25, "[a-zA-Z0-9]+", "admin_id"); e != nil {
		return e
	}

	// if e := util.ValidateRequireAndLengthAndRegex(p.Mode, true, 1, 25, "[a-zA-Z ]+", "mode"); e != nil {
	// 	return e
	// }

	// validating cost field with required, min length 0, max length 0 and regex check
	if p.Amount < 0 {
		return errors.New("amount is negative")
	}

	return nil
}

type EmployeeTransactionReport struct {
	EmployeeID string `json:"employee_id,required" bson:"employee_id,required"`
	Name       string `json:"name,required" bson:"name,required"`
	Amount     int64  `json:"amount,required,string" bson:"amount,required"`
}
