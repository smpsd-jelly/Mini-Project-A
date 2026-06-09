package postgres

import (
	"mini-project-a/internal/transactions/core/entity"

	"github.com/jmoiron/sqlx"
)

type TransactionsPostgresRepository struct {
	db *sqlx.DB
}

func NewTransactionsPostgresRepository(db *sqlx.DB) *TransactionsPostgresRepository {
	return &TransactionsPostgresRepository{db: db}
}

func (r *TransactionsPostgresRepository) CreateTransaction(transaction *entity.Transactions) (*entity.Transactions, error) {
	model := FromEntity(transaction)
	query := `INSERT INTO transactions (account_id, transaction_type, amount, balance_before, balance_after, description)
	VALUES (:account_id, :transaction_type, :amount, :balance_before, :balance_after, :description) RETURNING *`
	rows, err := r.db.NamedQuery(query, model)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var output Transactions
	if rows.Next() {
		err = rows.StructScan(&output)
		if err != nil {
			return nil, err
		}
	}
	result := output.ToEntity()
	return result, nil
}

func (r *TransactionsPostgresRepository) GetTransactionList() ([]*entity.Transactions, error) {
	query := `SELECT id, account_id, transaction_type, amount, balance_before, balance_after, description, created_at FROM transactions ORDER BY id DESC`
	var outputs []Transactions

	err := r.db.Select(&outputs, query)
	if err != nil {
		return nil, err
	}

	transcations := make([]*entity.Transactions, 0, len(outputs))
	for _, output := range outputs {
		transcations = append(transcations, output.ToEntity())
	}
	return transcations, nil

}
