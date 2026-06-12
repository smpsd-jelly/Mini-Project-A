package port

import (
	"mini-project-a/internal/accounts/core/entity"

	"github.com/jmoiron/sqlx"
)

type AccountsRepository interface {
	CreateAccount(account *entity.Accounts) (*entity.Accounts, error)
	GetAccountDetailByAccountNumber(accountNumber string) (*entity.Accounts, error)
	GetAccountList() ([]*entity.Accounts, error)
	GetAccountByCitizenID(citizenID string) (*entity.Accounts, error)
	CloseAccount(accountNumber string) (*entity.Accounts, error)
	UpdateBalanceTx(tx *sqlx.Tx, accountID int64, balance float64) error
}

type InitialDepositPort interface {
	CreateInitialDeposit(accountID int64, amount float64) error
}
