package postgres

import (
	"database/sql"
	"mini-project-a/internal/accounts/core/entity"
)

type AccountPostgresRepository struct {
	db *sql.DB
}

func NewAccountPostgresRepository(db *sql.DB) *AccountPostgresRepository {
	return &AccountPostgresRepository{
		db: db,
	}
}

func (r *AccountPostgresRepository) CreateAccount(account *entity.Accounts) error {
	model := &Accounts{
		ID:          account.ID,
		OwnerName:   account.OwnerName,
		CitizenID:   account.CitizenID,
		PhoneNumber: account.PhoneNumber,
		AccountType: account.AccountType,
		Balance:     account.Balance,
		Status:      account.Status,
		CreatedAt:   account.CreatedAt,
		UpdatedAt:   account.UpdatedAt,
	}
	query := `INSERT INTO accounts (owner_name, citizen_id, phone_number, account_type, balance, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, account_number`
	err := r.db.QueryRow(query,
		model.OwnerName,
		model.CitizenID,
		model.PhoneNumber,
		model.AccountType,
		model.Balance,
		model.Status,
		model.CreatedAt,
		model.UpdatedAt,
	).Scan(&model.ID, &model.AccountNumber)
	if err != nil {
		return err
	}
	account.ID = model.ID
	account.AccountNumber = model.AccountNumber
	return nil
}
