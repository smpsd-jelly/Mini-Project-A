package core

import (
	"mini-project-a/internal/shared/constant"
	"mini-project-a/internal/transactions/core/entity"
)

func (s *TransactionsService) CreateInitialDeposit(accountID int64, amount float64) error {
	description := "Initial deposit"

	transaction := &entity.Transactions{
		AccountID:       accountID,
		TransactionType: constant.TRANSACTIONS_TYPE_DEPOSIT,
		Amount:          amount,
		BalanceBefore:   0,
		BalanceAfter:    amount,
		Description:     &description,
	}
	_, err := s.CreateTransaction(transaction)
	return err
}
