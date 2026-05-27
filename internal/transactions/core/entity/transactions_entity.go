package entity

import "time"

type Transactions struct {
	ID              int64
	AccountID       int64
	TransactionType string
	Amount          float64
	BalanceBefore   float64
	BalanceAfter    float64
	Description     *string
	CreatedAt       time.Time
}
