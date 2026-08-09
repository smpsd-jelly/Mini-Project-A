package core

import (
	"database/sql"
	"errors"
	"mini-project-a/internal/accounts/core/entity"
)

func (s *AccountsService) GetAccountDetailByAccountNumber(accountNumber string) (*entity.Accounts, error) {
	account, err := s.accountsRepo.GetAccountDetailByAccountNumberRepo(accountNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAccountNotFound
		}

		return nil, err
	}

	return account, nil
}
