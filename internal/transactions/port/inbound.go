package port

import "mini-project-a/internal/transactions/core/entity"

type TransactionsService interface {
	CreateTransaction(transaction *entity.Transactions) error
}
