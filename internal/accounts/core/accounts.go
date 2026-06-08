package core

import (
	accountsPort "mini-project-a/internal/accounts/port"
)

type AccountsService struct {
	accountsRepo       accountsPort.AccountsRepository
	initialDepositPort accountsPort.InitialDepositPort
}

func NewAccountsService(accountsRepo accountsPort.AccountsRepository, initialDepositPort accountsPort.InitialDepositPort) *AccountsService {
	return &AccountsService{
		accountsRepo:       accountsRepo,
		initialDepositPort: initialDepositPort,
	}
}
