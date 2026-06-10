package http

import (
	"mini-project-a/internal/transactions/core/entity"
)

type TransactionRequest struct {
	Amount      float64 `json:"amount",gt=0"`
	Description *string `json:"description,omitempty"`
}

func (r *TransactionRequest) ToTransactionEntity() *entity.Transactions {
	return &entity.Transactions{
		Amount:      r.Amount,
		Description: r.Description,
	}
}
