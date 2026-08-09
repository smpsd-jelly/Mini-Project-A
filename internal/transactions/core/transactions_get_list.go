package core

import (
	"database/sql"
	"errors"
	"mini-project-a/internal/transactions/core/entity"
)

func (s *TransactionsService) GetTransactionList(
	accountNumber string,
) ([]*entity.Transactions, error) {

	_, err := s.accountReaderPort.GetAccountDetailByAccountNumberRepo(accountNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAccountNotFound

		}
		return nil, err
	}

	return s.transactionsRepo.GetTransactionList(accountNumber)
}
