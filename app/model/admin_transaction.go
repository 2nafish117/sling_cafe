package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sling_cafe/util"
	"time"
)

// AdminTransaction model
type AdminTransaction struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TransactionID   string             `json:"transaction_id,required" bson:"transaction_id,required"`     // autogenerated by server
	DateTime        time.Time          `json:"date_time,required" bson:"date_time,required"`               // time is generated server side using time.Now()
	TransactionType TransactionType    `json:"transaction_type,required" bson:"transaction_type,required"` // employee -> admin (employee payment) or admin -> manager (deposit)
	ManagerID       string             `json:"manager_id,omitempty" bson:"manager_id,omitempty"`           // will be there only for transactions that involve admin -> manager
	AdminID         string             `json:"admin_id,required" bson:"admin_id,required"`
	EmployeeID      string             `json:"employee_id,omitempty" bson:"employee_id,omitempty"` // will be there only for transactions that involve employee -> admin
	// VoucherNo       string             `json:"voucher_no,omitempty" bson:"voucher_no,omitempty"` // what is this ??
	Mode    Mode   `json:"mode,required" bson:"mode,required"`
	Amount  int64  `json:"amount,required,string" bson:"amount,required"`
	Remarks string `json:"remarks,omitempty" bson:"remarks,omitempty"`
}

// Validate fields
// This function validates meal data
// and return error is any
// all errors are related to the fields
func (p *AdminTransaction) Validate() error {
	// validating uid field with retuired, min length 1, max length 25 and regex check
	// if e := util.ValidateRequireAndLengthAndRegex(p.TransactionID, true, 1, 25, "[a-zA-Z0-9]+", "transaction_id"); e != nil {
	// 	return e
	// }

	// validating uid field with retuired, min length 1, max length 200 and regex check
	if e := util.ValidateRequireAndLengthAndRegex(p.AdminID, true, 1, 25, "[a-zA-Z0-9]+", "admin_id"); e != nil {
		return e
	}

	// validating uid field with retuired, min length 1, max length 200 and regex check
	// if e := util.ValidateRequireAndLengthAndRegex(p.ManagerID, false, 1, 25, "[a-zA-Z0-9]+", "manager_id"); e != nil {
	// 	return e
	// }

	// if e := util.ValidateRequireAndLengthAndRegex(p.Mode, true, 1, 25, "[a-zA-Z ]+", "mode"); e != nil {
	// 	return e
	// }

	// validating cost field with required, min length 0, max length 0 and regex check
	if p.Amount < 0 {
		return errors.New("amount is negative")
	}

	return nil
}

type AdminTransactionReport struct {
	AdminID string `json:"admin_id,required" bson:"admin_id,required"`
	Name    string `json:"name,required" bson:"name,required"`
	Amount  int64  `json:"amount,required,string" bson:"amount,required"`
}
