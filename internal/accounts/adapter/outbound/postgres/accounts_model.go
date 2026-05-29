package postgres

import (
	"mini-project-a/internal/accounts/core/entity"
	"time"
)

type Accounts struct {
	ID            int64      `db:"id"`
	AccountNumber string     `db:"account_number"`
	OwnerName     string     `db:"owner_name"`
	CitizenID     string     `db:"citizen_id"`
	PhoneNumber   string     `db:"phone_number"`
	AccountType   string     `db:"account_type"`
	Balance       float64    `db:"balance"`
	Status        string     `db:"status"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
}

type CreateAccountModel struct {
	OwnerName   string  `db:"owner_name"`
	CitizenID   string  `db:"citizen_id"`
	PhoneNumber string  `db:"phone_number"`
	AccountType string  `db:"account_type"`
	Balance     float64 `db:"balance"`
	Status      string  `db:"status"`
}

func (a *Accounts) ToEntity() *entity.Accounts {
	return &entity.Accounts{
		ID:            a.ID,
		AccountNumber: a.AccountNumber,
		OwnerName:     a.OwnerName,
		CitizenID:     a.CitizenID,
		PhoneNumber:   a.PhoneNumber,
		AccountType:   a.AccountType,
		Balance:       a.Balance,
		Status:        a.Status,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func FromEntity(e *entity.Accounts) *Accounts {
	return &Accounts{
		ID:            e.ID,
		AccountNumber: e.AccountNumber,
		OwnerName:     e.OwnerName,
		CitizenID:     e.CitizenID,
		PhoneNumber:   e.PhoneNumber,
		AccountType:   e.AccountType,
		Balance:       e.Balance,
		Status:        e.Status,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
}
