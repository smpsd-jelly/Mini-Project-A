package core

import (
	transactionPort "mini-project-a/internal/transactions/port"

	"github.com/jmoiron/sqlx"
)

type TransactionsService struct {
	db                *sqlx.DB
	transactionsRepo  transactionPort.TransactionsRepository
	accountReaderPort transactionPort.AccountReaderPort
}

func NewTransactionsService(
	db *sqlx.DB,
	transactionsRepo transactionPort.TransactionsRepository,
	accountReaderPort transactionPort.AccountReaderPort,
) *TransactionsService {
	return &TransactionsService{
		db:                db,
		transactionsRepo:  transactionsRepo,
		accountReaderPort: accountReaderPort,
	}
}
