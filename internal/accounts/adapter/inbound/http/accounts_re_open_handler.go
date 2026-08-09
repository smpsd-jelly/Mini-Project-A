package http

import (
	"errors"
	"mini-project-a/internal/accounts/core"
	responses "mini-project-a/internal/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AccountsHandler) ReopenAccountHandler(c *gin.Context) {
	var uri AccountNumberURI
	if err := c.ShouldBindUri(&uri); err != nil {

		responses.ValidateError(
			c,
			http.StatusBadRequest,
			"E1001",
			"Validation failed",
			responses.FormatValidationErrors(err),
		)
		return
	}
	accounts, err := h.accountsService.ReopenAccount(uri.AccountNumber)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			responses.Error(
				c,
				http.StatusNotFound,
				"E4001",
				"Account not found",
			)
		case errors.Is(err, core.ErrAccountAlreadyActive):
			responses.Error(
				c,
				http.StatusConflict,
				"E3001",
				"Account is already active",
			)
		default:
			responses.Error(
				c,
				http.StatusInternalServerError,
				"E9001",
				"Internal server error",
			)
		}
		return
	}

	account := ToReOpenAccountDetailResponse(accounts)

	responses.Success(
		c,
		http.StatusOK,
		"S0000",
		"Account Re-open successfully",
		account,
	)

}
