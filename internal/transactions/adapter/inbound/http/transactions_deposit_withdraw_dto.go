package http

import (
	"mini-project-a/internal/transactions/core/entity"
)

type AccountNumberURI struct {
	AccountNumber string `uri:"accountNumber" binding:"required,len=10,numeric"`
}

type TransactionRequest struct {
	Amount      float64 `json:"amount" binding:"required,numeric,gt=0"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=255"`
}

func (r *TransactionRequest) ToTransactionEntity() *entity.Transactions {
	return &entity.Transactions{
		Amount:      r.Amount,
		Description: r.Description,
	}
}
