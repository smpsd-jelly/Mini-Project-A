package core

import (
	"errors"
	"mini-project-a/internal/accounts/core/entity"
	"mini-project-a/internal/shared/constant"
)

func (s *AccountsService) CreateAccount(account *entity.Accounts) (*entity.Accounts, error) {

	existingAccount, err := s.accountsRepo.GetAccountByCitizenID(account.CitizenID)
	if err != nil {
		return nil, err
	}
	if existingAccount != nil {
		return nil, errors.New("account with the same citizen ID already exists")
	}

	account.Status = constant.ACCOUNTS_STATUS_ACTIVE

	result, err := s.accountsRepo.CreateAccount(account)
	if err != nil {
		return nil, err
	}

	if account.Balance > 0 {
		err = s.initialDepositPort.CreateInitialDeposit(result.ID, account.Balance)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
