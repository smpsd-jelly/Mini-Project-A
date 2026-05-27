package port

import "mini-project-a/internal/accounts/core/entity"

type AccountsRepository interface {
	CreateAccount(account *entity.Accounts) error
}
