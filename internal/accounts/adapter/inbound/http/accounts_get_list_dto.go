package http

import "mini-project-a/internal/accounts/core/entity"

type GetAccountListResponse struct {
	Items []GetAccountDetailResponse `json:"items"`
}

func ToGetAccountListResponse(accounts []*entity.Accounts) GetAccountListResponse {
	items := make([]GetAccountDetailResponse, 0, len(accounts))

	for _, account := range accounts {
		items = append(items, ToGetAccountDetailResponse(account))
	}

	return GetAccountListResponse{
		Items: items,
	}
}
