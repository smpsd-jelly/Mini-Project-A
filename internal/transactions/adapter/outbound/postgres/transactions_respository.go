package postgres

import (
	"database/sql"
	"mini-project-a/internal/transactions/core/entity"
)

type TransactionsPostgresRepository struct {
	db *sql.DB
}

func NewTransactionsPostgresRepository(db *sql.DB) *TransactionsPostgresRepository {
	return &TransactionsPostgresRepository{db: db}
}

func (r *TransactionsPostgresRepository) CreateTransaction(transaction *entity.Transactions) error {
	model := FromEntity(transaction)
	query := `INSERT INTO transactions (account_id, transaction_type, amount, balance_before, balance_after, description, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := r.db.QueryRow(query,
		model.AccountID,
		model.TransactionType,
		model.Amount,
		model.BalanceBefore,
		model.BalanceAfter,
		model.Description,
		model.CreatedAt,
	).Scan(&model.ID)
	if err != nil {
		return err
	}
	transaction.ID = model.ID
	return nil
}
