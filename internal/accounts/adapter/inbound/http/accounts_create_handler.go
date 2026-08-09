package http

import (
	"errors"
	"mini-project-a/internal/accounts/core"
	"mini-project-a/internal/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AccountsHandler) CreateAccount(c *gin.Context) {
	var req CreateAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateError(
			c,
			http.StatusBadRequest,
			"E1001",
			"Validation failed.",
			response.FormatValidationErrors(err),
		)

		return
	}

	account := req.ToAccountEntity()

	createdAccount, err := h.accountsService.CreateAccount(account)

	if err != nil {

		switch {
		case errors.Is(err, core.ErrCitizenIDAlreadyExists):
			response.Error(
				c,
				http.StatusConflict,
				"E3001",
				"Account with the same citizen ID already exists",
			)

		case errors.Is(err, core.ErrAccountNumberExists):
			response.Error(
				c,
				http.StatusConflict,
				"E3001",
				"Account Number already exists",
			)

		default:
			response.Error(
				c,
				http.StatusInternalServerError,
				"E9001",
				"Internal server error",
			)
			return
		}
	}

	response.Success(
		c,
		http.StatusCreated,
		"S0001",
		"Account created successfully.",
		createdAccount,
	)
}
