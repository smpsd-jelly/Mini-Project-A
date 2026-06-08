package http

import "mini-project-a/internal/accounts/core/entity"

type GetAccountDetailResponse struct {
	AccountNumber string  `json:"account_number"`
	OwnerName     string  `json:"owner_name"`
	AccountType   string  `json:"account_type"`
	Balance       float64 `json:"balance"`
	Status        string  `json:"status"`
}

func ToGetAccountDetailResponse(account *entity.Accounts) GetAccountDetailResponse {
	return GetAccountDetailResponse{
		AccountNumber: account.AccountNumber,
		OwnerName:     account.OwnerName,
		AccountType:   account.AccountType,
		Balance:       account.Balance,
		Status:        account.Status,
	}
}
