package core

import (
	"mini-project-a/internal/transactions/core/entity"
	"mini-project-a/internal/transactions/port"
)

type TransactionsService struct {
	transactionsRepo port.TransactionsRepository
}

func NewTransactionsService(transactionsRepo port.TransactionsRepository) *TransactionsService {
	return &TransactionsService{
		transactionsRepo: transactionsRepo,
	}
}

func (s *TransactionsService) CreateTransaction(transaction *entity.Transactions) error {
	return s.transactionsRepo.CreateTransaction(transaction)
}
