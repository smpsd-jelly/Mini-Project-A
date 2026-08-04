package http

import "mini-project-a/internal/accounts/core/entity"

func ToReOpenAccountDetailResponse(account *entity.Accounts) GetAccountDetailResponse {
	return GetAccountDetailResponse{
		AccountNumber: account.AccountNumber,
		OwnerName:     account.OwnerName,
		AccountType:   account.AccountType,
		Balance:       account.Balance,
		Status:        account.Status,
	}
}
