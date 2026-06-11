package port

import (
	accountEntity "mini-project-a/internal/accounts/core/entity"
	"mini-project-a/internal/transactions/core/entity"

	"github.com/jmoiron/sqlx"
)

type TransactionsRepository interface {
	CreateTransaction(transaction *entity.Transactions) (*entity.Transactions, error)
	BeginTx() (*sqlx.Tx, error)
	CreateTransactionTx(tx *sqlx.Tx, transaction *entity.Transactions) (*entity.Transactions, error)
}

type AccountReaderPort interface {
	GetAccountDetailByAccountNumber(
		accountNumber string,
	) (*accountEntity.Accounts, error)
	UpdateBalanceTx(tx *sqlx.Tx, accountID int64, balance float64) error
}
