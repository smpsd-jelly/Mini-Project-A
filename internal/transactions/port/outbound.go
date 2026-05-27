package port

import "mini-project-a/internal/transactions/core/entity"

type TransactionsRepository interface {
	CreateTransaction(transaction *entity.Transactions) error
}
