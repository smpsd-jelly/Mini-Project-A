package core

import (
	"errors"
	"mini-project-a/internal/accounts/core/entity"
	"mini-project-a/internal/shared/constant"
)

func (s *AccountsService) ReopenAccount(accountNumber string) (*entity.Accounts, error) {
	account, err := s.accountsRepo.GetAccountDetailByAccountNumberRepo(accountNumber)
	if err != nil {
		return nil, err
	}

	if account.Status == constant.ACCOUNTS_STATUS_ACTIVE {
		return nil, errors.New("This Account is Already ACTIVE")
	}

	return s.accountsRepo.ReopenAccountRepo(accountNumber)
}
