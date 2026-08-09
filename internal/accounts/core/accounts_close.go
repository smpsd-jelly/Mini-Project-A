package core

import (
	"database/sql"
	"errors"
	"mini-project-a/internal/accounts/core/entity"
	"mini-project-a/internal/shared/constant"
)

func (s *AccountsService) CloseAccount(accountNumber string) (*entity.Accounts, error) {
	account, err := s.accountsRepo.GetAccountDetailByAccountNumberRepo(accountNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAccountNotFound

		}
		return nil, err
	}

	if account.Status == constant.ACCOUNTS_STATUS_CLOSED {
		return nil, ErrAccountAlreadyClosed
	}

	return s.accountsRepo.CloseAccountRepo(accountNumber)
}
