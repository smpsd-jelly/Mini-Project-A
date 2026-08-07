package port

import (
	"mini-project-a/internal/accounts/core/entity"
	sharedPort "mini-project-a/internal/shared/port"
)

type Tx interface {
	Commit() error
	Rollback() error
}

type AccountsRepository interface {
	CreateAccountRepo(account *entity.Accounts) (*entity.Accounts, error)
	GetAccountDetailByAccountNumberRepo(accountNumber string) (*entity.Accounts, error)
	GetAccountListRepo() ([]*entity.Accounts, error)
	GetAccountByCitizenIDRepo(citizenID string) (*entity.Accounts, error)
	CloseAccountRepo(accountNumber string) (*entity.Accounts, error)
	UpdateBalanceTx(tx sharedPort.Tx, accountID int64, balance float64) error
	ReopenAccountRepo(accountNumber string) (*entity.Accounts, error)
}

type InitialDepositPort interface {
	CreateInitialDeposit(accountID int64, amount float64) error
}
