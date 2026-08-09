package port

import (
	accountEntity "mini-project-a/internal/accounts/core/entity"
	sharedPort "mini-project-a/internal/shared/port"
	"mini-project-a/internal/transactions/core/entity"
)

type TransactionsRepository interface {
	CreateTransaction(transaction *entity.Transactions) (*entity.Transactions, error)
	BeginTx() (sharedPort.Tx, error)
	CreateTransactionTx(tx sharedPort.Tx, transaction *entity.Transactions) (*entity.Transactions, error)
	GetTransactionList(accountNumber string) ([]*entity.Transactions, error)
}

type AccountReaderPort interface {
	GetAccountDetailByAccountNumberRepo(
		accountNumber string,
	) (*accountEntity.Accounts, error)
	UpdateBalanceTx(tx sharedPort.Tx, accountID int64, balance float64) error
}
