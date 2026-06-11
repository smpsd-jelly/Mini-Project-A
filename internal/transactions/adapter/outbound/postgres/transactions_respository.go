package postgres

import (
	"mini-project-a/internal/transactions/core/entity"

	"github.com/jmoiron/sqlx"
)

type namedQueryer interface {
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
}

type TransactionsPostgresRepository struct {
	db *sqlx.DB
}

func NewTransactionsPostgresRepository(db *sqlx.DB) *TransactionsPostgresRepository {
	return &TransactionsPostgresRepository{db: db}
}

func (r *TransactionsPostgresRepository) BeginTx() (*sqlx.Tx, error) {
	return r.db.Beginx()
}

func (r *TransactionsPostgresRepository) CreateTransaction(
	transaction *entity.Transactions,
) (*entity.Transactions, error) {
	return r.createTransaction(r.db, transaction)
}

func (r *TransactionsPostgresRepository) CreateTransactionTx(
	tx *sqlx.Tx,
	transaction *entity.Transactions,
) (*entity.Transactions, error) {
	return r.createTransaction(tx, transaction)
}

func (r *TransactionsPostgresRepository) createTransaction(
	db namedQueryer,
	transaction *entity.Transactions,
) (*entity.Transactions, error) {
	model := FromEntity(transaction)

	query := `
		INSERT INTO transactions (
			account_id,
			transaction_type,
			amount,
			balance_before,
			balance_after,
			description
		)
		VALUES (
			:account_id,
			:transaction_type,
			:amount,
			:balance_before,
			:balance_after,
			:description
		)
		RETURNING *
	`

	rows, err := db.NamedQuery(query, model)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var output Transactions
	if rows.Next() {
		if err := rows.StructScan(&output); err != nil {
			return nil, err
		}
	}

	return output.ToEntity(), nil
}
