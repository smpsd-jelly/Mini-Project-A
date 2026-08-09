// package core

// import "errors"

// var (

// 	//citizen id
// 	ErrCitizenIDAlreadyExists = errors.New("citizen id already exists")
// 	ErrCitizenIDMustBe13      = errors.New("citizen id must be 13 digits")
// 	ErrCitizenRequire         = errors.New("citizen id is required")

// 	//account
// 	ErrAccountNotFound      = errors.New("account not found")
// 	ErrAccountAlreadyClosed = errors.New("account already closed")
// 	ErrAccountAlreadyActive = errors.New("account already active")

// 	//account type
// 	ErrAccountType    = errors.New("account type must be SAVING or CURRENT")
// 	ErrAccountRequire = errors.New("account type is required")

// 	//account number
// 	ErrAccountNumberExists    = errors.New("account number already exists")
// 	ErrAccountNumberFormat    = errors.New("account number format is invalid")
// 	ErrAccountNumberMusetBe10 = errors.New("account number must be 10 digits")
// 	ErrAccountNumberRequire   = errors.New("account number is required")

// 	//owner name
// 	ErrOwnerNameRequire = errors.New("owner name is required")

// 	//phone name
// 	ErrPhoneNumberRequire = errors.New("phone number is required")

// 	//amount
// 	ErrAmountGreaterThan0 = errors.New("amount must be greater than 0")
// 	ErrAmountIsNotNumber  = errors.New("amount must be a number")
// 	ErrAmountRequire      = errors.New("amount is required")

// 	//description
// 	Errdescription255 = errors.New("description must not exceed 255 characters")

// 	//balance
// 	ErrInsufficientBalance = errors.New("insufficient balance") //กรณียอดเงินไม่เพียงพอ

// 	ErrNilBody = errors.New("invalid request body")
// )

package core

import "errors"

var (
	ErrCitizenIDAlreadyExists = errors.New("citizen id already exists")
	ErrAccountNumberExists    = errors.New("account number already exists")

	ErrAccountNotFound      = errors.New("account not found")
	ErrAccountAlreadyClosed = errors.New("account already closed")
	ErrAccountAlreadyActive = errors.New("account already active")

	ErrInsufficientBalance = errors.New("insufficient balance")
)
//bussines errors