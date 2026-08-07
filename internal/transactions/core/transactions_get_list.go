package core

import (
	"errors"
	"mini-project-a/internal/transactions/core/entity"
)

func (s *TransactionsService) GetTransactionList(
	accountNumber string,
) ([]*entity.Transactions, error) {
	account, err := s.accountReaderPort.GetAccountDetailByAccountNumberRepo(accountNumber)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.New("account not found")
	}

	return s.transactionsRepo.GetTransactionList(accountNumber)
}
