package core

import "mini-project-a/internal/accounts/core/entity"

func (s *AccountsService) GetAccountList() ([]*entity.Accounts, error) {
	return s.accountsRepo.GetAccountListRepo()
}
