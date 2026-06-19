package core

import "mini-project-a/internal/transactions/core/entity"

func (s *TransactionsService) CreateTransaction(transaction *entity.Transactions) (*entity.Transactions, error) {
	return s.transactionsRepo.CreateTransaction(transaction)
}
