package core

import (
	"mini-project-a/internal/accounts/core/entity"
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

func (s *AccountsService) CreateAccount(account *entity.Accounts) (*entity.Accounts, error) {
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

func (s *AccountsService) GetAccountDetailByAccountNumber(accountNumber string) (*entity.Accounts, error) {
	return s.accountsRepo.GetAccountDetailByAccountNumber(accountNumber)
}
