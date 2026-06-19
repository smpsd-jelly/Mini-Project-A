package http

import (
	"mini-project-a/internal/transactions/core"
	"mini-project-a/internal/transactions/core/entity"
)

type TransactionsHandler struct {
	transactionsService *core.TransactionsService
}

func NewTransactionsHandler(transactionsService *core.TransactionsService) *TransactionsHandler {
	return &TransactionsHandler{
		transactionsService: transactionsService,
	}
}

type GetTransactionDetailResponse struct {
	ID              int64   `json:"id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	BalanceBefore   float64 `json:"balance_before"`
	BalanceAfter    float64 `json:"balance_after"`
	Description     *string `json:"description"`
}

func ToGetTransactionDetailResponse(
	transaction *entity.Transactions,
) GetTransactionDetailResponse {
	return GetTransactionDetailResponse{
		ID:              transaction.ID,
		TransactionType: transaction.TransactionType,
		Amount:          transaction.Amount,
		BalanceBefore:   transaction.BalanceBefore,
		BalanceAfter:    transaction.BalanceAfter,
		Description:     transaction.Description,
	}
}
