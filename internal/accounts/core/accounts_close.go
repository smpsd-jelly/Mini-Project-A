package core

import (
	"errors"
	"mini-project-a/internal/accounts/core/entity"
	"mini-project-a/internal/shared/constant"
)

func (s *AccountsService) CloseAccount(accountNumber string) (*entity.Accounts, error) {
	account, err := s.accountsRepo.GetAccountDetailByAccountNumberRepo(accountNumber)
	if err != nil {
		return nil, err
	}

	if account.Status == constant.ACCOUNTS_STATUS_CLOSED {
		return nil, errors.New("account is already closed")
	}

	return s.accountsRepo.CloseAccountRepo(accountNumber)
}
