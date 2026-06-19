package http

import "mini-project-a/internal/transactions/core"

type TransactionsHandler struct {
	transactionsService *core.TransactionsService
}

func NewTransactionsHandler(transactionsService *core.TransactionsService) *TransactionsHandler {
	return &TransactionsHandler{
		transactionsService: transactionsService,
	}
}
