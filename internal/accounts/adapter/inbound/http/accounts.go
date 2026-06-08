package http

import (
	"mini-project-a/internal/accounts/core"
)

type AccountsHandler struct {
	accountsService *core.AccountsService
}

func NewAccountsHandler(accountsService *core.AccountsService) *AccountsHandler {
	return &AccountsHandler{
		accountsService: accountsService,
	}
}
