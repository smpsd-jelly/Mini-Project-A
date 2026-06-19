package port

import (
	"mini-project-a/internal/transactions/core/entity"
)

type TransactionsService interface {
	CreateTransaction(transaction *entity.Transactions) (*entity.Transactions, error)
	CreateInitialDeposit(accountID int64, amount float64) error
	Deposit(accountNumber string, transaction *entity.Transactions) (*entity.Transactions, error)
	Withdraw(accountNumber string, transaction *entity.Transactions) (*entity.Transactions, error)
}
