package port

import "mini-project-a/internal/accounts/core/entity"

type AccountsService interface {
	CreateAccount(account *entity.Accounts) error
}
