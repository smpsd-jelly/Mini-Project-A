package core

import "errors"

var (
	ErrAccountNumberExists = errors.New("account number already exists")

	ErrAccountNotFound = errors.New("account not found")
	ErrAccountIsClosed = errors.New("account is closed")
	ErrAccountIsActive = errors.New("account is active")

	ErrInsufficientBalance = errors.New("insufficient balance")
)
