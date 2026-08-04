package port

import "mini-project-a/internal/accounts/core/entity"

type AccountsService interface {
	CreateAccount(account *entity.Accounts) (*entity.Accounts, error)
	GetAccountDetailByAccountNumber(accountNumber string) (*entity.Accounts, error)
	GetAccountList() ([]*entity.Accounts, error)
	CloseAccount(accountNumber string) (*entity.Accounts, error)
	ReopenAccount(accountNumber string) (*entity.Accounts, error)
}
