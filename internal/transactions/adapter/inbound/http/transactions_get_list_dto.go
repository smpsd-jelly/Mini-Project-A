package http

import (
	"mini-project-a/internal/transactions/core/entity"
)

type GetTransactionListResponse struct {
	Items []GetTransactionDetailResponse `json:"items"`
}

func ToGetTransactionListResponse(transactions []*entity.Transactions) GetTransactionListResponse {
	items := make([]GetTransactionDetailResponse, 0, len(transactions))

	for _, transaction := range transactions {
		items = append(items, ToGetTransactionDetailResponse(transaction))
	}

	return GetTransactionListResponse{
		Items: items,
	}
}
