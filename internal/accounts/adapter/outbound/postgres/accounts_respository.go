package postgres

import (
	"mini-project-a/internal/accounts/core/entity"

	"github.com/jmoiron/sqlx"
)

type AccountPostgresRepository struct {
	db *sqlx.DB
}

func NewAccountPostgresRepository(db *sqlx.DB) *AccountPostgresRepository {
	return &AccountPostgresRepository{
		db: db,
	}
}

func (r *AccountPostgresRepository) CreateAccount(account *entity.Accounts) (*entity.Accounts, error) {
	model := FromEntity(account)
	query := `INSERT INTO accounts (owner_name, citizen_id, phone_number, account_type, balance, status)
	VALUES (
			:owner_name,
			:citizen_id,
			:phone_number,
			:account_type,
			:balance,
			:status) 
		RETURNING *`
	rows, err := r.db.NamedQuery(query, model)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var output Accounts
	if rows.Next() {
		err = rows.StructScan(&output)
		if err != nil {
			return nil, err
		}
	}
	result := output.ToEntity()
	return result, nil
}

func (r *AccountPostgresRepository) GetAccountDetailByAccountNumber(accountNumber string) (*entity.Accounts, error) {
	query := `
		SELECT 
			account_number,
			owner_name,
			account_type,
			balance,
			status
		FROM accounts
		WHERE account_number = $1
	`

	var output Accounts

	err := r.db.Get(&output, query, accountNumber)
	if err != nil {
		return nil, err
	}

	return output.ToEntity(), nil
}

func (r *AccountPostgresRepository) GetAccountList() ([]*entity.Accounts, error) {
	query := `
		SELECT 
			account_number,
			owner_name,
			account_type,
			balance,
			status
		FROM accounts
	`

	var outputs []Accounts

	err := r.db.Select(&outputs, query)
	if err != nil {
		return nil, err
	}

	accounts := make([]*entity.Accounts, 0, len(outputs))
	for _, output := range outputs {
		accounts = append(accounts, output.ToEntity())
	}

	return accounts, nil
}

func (r *AccountPostgresRepository) CloseAccount(accountNumber string) (*entity.Accounts, error) {
	query := `
        UPDATE accounts
        SET status = 'CLOSED',
		    updated_at = NOW()
        WHERE account_number = $1
        RETURNING *`

	var output Accounts

	err := r.db.Get(&output, query, accountNumber)
	if err != nil {
		return nil, err
	}

	return output.ToEntity(), nil
}
