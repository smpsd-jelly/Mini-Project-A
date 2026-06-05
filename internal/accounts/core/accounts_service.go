package core

import (
	"errors"
	"mini-project-a/internal/accounts/core/entity"
	accountsPort "mini-project-a/internal/accounts/port"
	"mini-project-a/internal/shared/constant"
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

func (s *AccountsService) GetAccountDetailByAccountNumber(accountNumber string) (*entity.Accounts, error) {
	return s.accountsRepo.GetAccountDetailByAccountNumber(accountNumber)
}

func (s *AccountsService) GetAccountList() ([]*entity.Accounts, error) {
	return s.accountsRepo.GetAccountList()
}

func (s *AccountsService) CloseAccount(accountNumber string) (*entity.Accounts, error) {
	account, err := s.accountsRepo.GetAccountDetailByAccountNumber(accountNumber)
	if err != nil {
		return nil, err
	}

	if account.Status == "CLOSED" {
		return nil, errors.New("account is already closed")
	}

	return s.accountsRepo.CloseAccount(accountNumber)
}
