package core

import (
	"mini-project-a/internal/accounts/core/entity"
	"mini-project-a/internal/accounts/port"
)

type AccountsService struct {
	accountsRepo port.AccountsRespository
}

func NewAccountsService(accountsRepo port.AccountsRespository) *AccountsService {
	return &AccountsService{
		accountsRepo: accountsRepo,
	}
}

func (s *AccountsService) CreateAccount(account *entity.Accounts) error {
	return s.accountsRepo.CreateAccount(account)
}
