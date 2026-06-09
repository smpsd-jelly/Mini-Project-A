package core

import "mini-project-a/internal/transactions/core/entity"

func (s *TransactionsService) GetTransactionList() ([]*entity.Transactions, error) {
	return s.transactionsRepo.GetTransactionList()
}
