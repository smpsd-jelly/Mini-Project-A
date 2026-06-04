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

func (s *TransactionsService) CreateTransaction(transaction *entity.Transactions) (*entity.Transactions, error) {
	return s.transactionsRepo.CreateTransaction(transaction)
}

func (s *TransactionsService) CreateInitialDeposit(accountID int64, amount float64) error {
	description := "Initial deposit"

	transaction := &entity.Transactions{
		AccountID:       accountID,
		TransactionType: "DEPOSIT",
		Amount:          amount,
		BalanceBefore:   0,
		BalanceAfter:    amount,
		Description:     &description,
	}
	_, err := s.CreateTransaction(transaction)
	return err
}
