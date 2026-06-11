package port

import (
	accountEntity "mini-project-a/internal/accounts/core/entity"
	"mini-project-a/internal/transactions/core/entity"
)

type TransactionsRepository interface {
	CreateTransaction(transaction *entity.Transactions) (*entity.Transactions, error)
	GetTransactionList(accountNumber string) ([]*entity.Transactions, error)
}

type AccountReaderPort interface {
	GetAccountDetailByAccountNumber(
		accountNumber string,
	) (*accountEntity.Accounts, error)
	UpdateBalance(accountID int64, balance float64) error
}
