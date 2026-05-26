package port

import "mini-project-a/internal/accounts/core/entity"

type AccountsRespository interface {
	CreateAccount(account *entity.Accounts) error
}
