package postgres

import (
	"database/sql"
	"mini-project-a/internal/accounts/core/entity"
	"mini-project-a/internal/shared/constant"
	sharedPort "mini-project-a/internal/shared/port"

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

func (r *AccountPostgresRepository) GetAccountByCitizenID(citizenID string) (*entity.Accounts, error) {
	var account Accounts

	query := `
		SELECT id, account_number, owner_name, citizen_id,
		       phone_number, account_type, balance,
		       status, created_at, updated_at
		FROM accounts
		WHERE citizen_id = $1
	`

	err := r.db.Get(&account, query, citizenID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return account.ToEntity(), nil
}

func (r *AccountPostgresRepository) GetAccountDetailByAccountNumber(accountNumber string) (*entity.Accounts, error) {
	query := `SELECT id, account_number, owner_name, citizen_id, phone_number, account_type, balance, status, created_at, updated_at FROM accounts WHERE account_number = $1`

	var output Accounts

	err := r.db.Get(&output, query, accountNumber)
	if err != nil {
		return nil, err
	}

	return output.ToEntity(), nil
}

func (r *AccountPostgresRepository) GetAccountList() ([]*entity.Accounts, error) {
	query := `SELECT id, account_number, owner_name, citizen_id, phone_number, account_type, balance, status, created_at, updated_at FROM accounts`

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
        SET status = $1,
		    updated_at = NOW()
        WHERE account_number = $2
        RETURNING *`

	var output Accounts

	err := r.db.Get(&output, query, constant.ACCOUNTS_STATUS_CLOSED, accountNumber)
	if err != nil {
		return nil, err
	}

	return output.ToEntity(), nil
}

func (r *AccountPostgresRepository) UpdateBalanceTx(
	tx sharedPort.Tx,
	accountID int64,
	balance float64,
) error {
	query := `
		UPDATE accounts
		SET balance = $1, updated_at = NOW()
		WHERE id = $2
	`

	_, err := tx.Exec(query, balance, accountID)
	return err
}
