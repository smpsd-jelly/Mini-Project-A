package postgres

import (
	"mini-project-a/internal/transactions/core/entity"
	"time"
)

type Transactions struct {
	ID              int64     `db:"id"`
	AccountID       int64     `db:"account_id"`
	TransactionType string    `db:"transaction_type"`
	Amount          float64   `db:"amount"`
	BalanceBefore   float64   `db:"balance_before"`
	BalanceAfter    float64   `db:"balance_after"`
	Description     string    `db:"description"`
	CreatedAt       time.Time `db:"created_at"`
}

func (t *Transactions) ToEntity() *entity.Transactions {
	return &entity.Transactions{
		ID:              t.ID,
		AccountID:       t.AccountID,
		TransactionType: t.TransactionType,
		Amount:          t.Amount,
		BalanceBefore:   t.BalanceBefore,
		BalanceAfter:    t.BalanceAfter,
		Description:     &t.Description,
		CreatedAt:       t.CreatedAt,
	}
}

func FromEntity(e *entity.Transactions) *Transactions {
	description := ""
	if e.Description != nil {
		description = *e.Description
	}
	return &Transactions{
		ID:              e.ID,
		AccountID:       e.AccountID,
		TransactionType: e.TransactionType,
		Amount:          e.Amount,
		BalanceBefore:   e.BalanceBefore,
		BalanceAfter:    e.BalanceAfter,
		Description:     description,
		CreatedAt:       e.CreatedAt,
	}
}
