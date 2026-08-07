package core

import "mini-project-a/internal/accounts/core/entity"

func (s *AccountsService) GetAccountDetailByAccountNumber(accountNumber string) (*entity.Accounts, error) {
	return s.accountsRepo.GetAccountDetailByAccountNumberRepo(accountNumber)
}
