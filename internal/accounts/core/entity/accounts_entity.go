package entity

import "time"

type Accounts struct {
	ID            int64
	AccountNumber string
	OwnerName     string
	CitizenID     string
	PhoneNumber   string
	AccountType   string
	Balance       float64
	Status        string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}
