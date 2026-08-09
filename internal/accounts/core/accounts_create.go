package core

import (
	"mini-project-a/internal/accounts/core/entity"
	"mini-project-a/internal/shared/constant"
)

func (s *AccountsService) CreateAccount(account *entity.Accounts) (*entity.Accounts, error) {

	existingAccount, err := s.accountsRepo.GetAccountByCitizenIDRepo(account.CitizenID)
	if err != nil {
		return nil, err
	}
	if existingAccount != nil {
		return nil, ErrCitizenIDAlreadyExists
	}

	existingAccountNumber, err := s.GetAccountDetailByAccountNumber(account.AccountNumber)
	if err != nil {
		return nil, err
	}
	if existingAccountNumber != nil {
		return nil, ErrAccountNumberExists
	}

	account.Status = constant.ACCOUNTS_STATUS_ACTIVE

	result, err := s.accountsRepo.CreateAccountRepo(account)
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
