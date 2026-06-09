package core

import (
	transactionPort "mini-project-a/internal/transactions/port"
)

type TransactionsService struct {
	transactionsRepo  transactionPort.TransactionsRepository
	accountReaderPort transactionPort.AccountReaderPort
}

func NewTransactionsService(
	transactionsRepo transactionPort.TransactionsRepository,
	accountReaderPort transactionPort.AccountReaderPort,
) *TransactionsService {
	return &TransactionsService{
		transactionsRepo:  transactionsRepo,
		accountReaderPort: accountReaderPort,
	}
}
