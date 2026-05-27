package core

import (
	"mini-project-a/internal/accounts/core/entity"
	accountsPort "mini-project-a/internal/accounts/port"
	transactionsCore "mini-project-a/internal/transactions/core"
	transactionEntity "mini-project-a/internal/transactions/core/entity"
	"time"
)

type AccountsService struct {
	accountsRepo        accountsPort.AccountsRepository
	transactionsService *transactionsCore.TransactionsService
}

func NewAccountsService(accountsRepo accountsPort.AccountsRepository, transactionsService *transactionsCore.TransactionsService) *AccountsService {
	return &AccountsService{
		accountsRepo:        accountsRepo,
		transactionsService: transactionsService,
	}
}

func (s *AccountsService) CreateAccount(account *entity.Accounts) error {
	err := s.accountsRepo.CreateAccount(account)
	if err != nil {
		return err
	}

	if account.Balance > 0 {
		description := "Initial deposit"

		transaction := &transactionEntity.Transactions{
			AccountID:       account.ID,
			TransactionType: "DEPOSIT",
			Amount:          account.Balance,
			BalanceBefore:   0,
			BalanceAfter:    account.Balance,
			Description:     &description,
			CreatedAt:       time.Now(),
		}

		err = s.transactionsService.CreateTransaction(transaction)
		if err != nil {
			return err
		}
	}

	return nil
}
