package port

import "mini-project-a/internal/accounts/core/entity"

type AccountsRepository interface {
	CreateAccount(account *entity.Accounts) (*entity.Accounts, error)
	GetAccountDetailByAccountNumber(accountNumber string) (*entity.Accounts, error)
	GetAccountList() ([]*entity.Accounts, error)
	GetAccountByCitizenID(citizenID string) (*entity.Accounts, error)
	CloseAccount(accountNumber string) (*entity.Accounts, error)
}

type InitialDepositPort interface {
	CreateInitialDeposit(accountID int64, amount float64) error
}
