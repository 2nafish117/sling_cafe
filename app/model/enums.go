package model

// transaction type enum
type TransactionType = string

const (
	// EmployeePayment transaction type employee payment
	EmployeePayment TransactionType = "EmployeePayment"
	// AdminDeposit transaction type admin deposit
	AdminDeposit = "AdminDeposit"
	// MonthlyPosting is monthly aggregate type
	MonthlyPosting = "MonthlyPosting"
)

// MealID meal type enum
type MealID = string

const (
	// Breakfast meal type
	Breakfast MealID = "Breakfast"
	// Lunch meal type
	Lunch = "Lunch"
	// Snack meal type
	Snack = "Snack"
)

// Mode of payment type enum
type Mode = string

const (
	// Cash payment type
	Cash Mode = "Cash"
	// CreditCard payment type
	CreditCard = "CreditCard"
	// DebitCard payment type
	DebitCard = "DebitCard"
	// Upi payment type
	Upi = "Upi"
)

// type AdminTransactionType string

// const (
// 	// EmployeePayment transaction type employee payment
// 	EmployeePayment AdminTransactionType = "EmployeePayment"
// 	// AdminDeposit transaction type admin deposit
// 	AdminDeposit = "AdminDeposit"
// )

// type EmployeeTransactionType string

// const (
// 	// EmployeePayment transaction type employee payment
// 	EmployeePayment EmployeeTransactionType = "EmployeePayment"
// 	// MonthlyPosting is monthly aggregate type
// 	MonthlyPosting = "MonthlyPosting"
// )
