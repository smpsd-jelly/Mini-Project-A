package http

import (
	"mini-project-a/internal/transactions/core/entity"
)

type DepositCreateRequest struct {
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Description *string `json:"description,omitempty"`
}

func (r *DepositCreateRequest) ToTransactionEntity() *entity.Transactions {
	return &entity.Transactions{
		Amount:      r.Amount,
		Description: r.Description,
	}
}
