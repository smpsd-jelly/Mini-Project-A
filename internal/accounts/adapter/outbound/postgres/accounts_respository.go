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

func (r *AccountPostgresRepository) GetAccountByCitizenID(citizenID string) (*entity.Accounts, error) {
	var account Accounts
	query := `SELECT id, owner_name, citizen_id, phone_number, account_type, balance, status, created_at, updated_at FROM accounts WHERE citizen_id = $1`
	err := r.db.Get(&account, query, citizenID)
	if err != nil {
		return nil, err
	}
	result := account.ToEntity()
	return result, nil
}
