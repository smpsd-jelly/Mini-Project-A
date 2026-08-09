package http

import (
	"errors"
	"mini-project-a/internal/accounts/core"
	"mini-project-a/internal/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AccountsHandler) CloseAccount(c *gin.Context) {
	// accountNumber := c.Param("accountNumber")

	var uri AccountNumberURI

	if err := c.ShouldBindUri(&uri); err != nil {
		response.ValidateError(
			c,
			http.StatusBadRequest,
			"E1001",
			"Validation failed",
			response.FormatValidationErrors(err),
		)
		return
	}

	account, err := h.accountsService.CloseAccount(uri.AccountNumber)

	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			response.Error(
				c,
				http.StatusNotFound,
				"E4001",
				"Account not found",
			)

		case errors.Is(err, core.ErrAccountAlreadyClosed):
			response.Error(
				c,
				http.StatusConflict,
				"E3001",
				"Account is already closed",
			)

		default:
			response.Error(
				c,
				http.StatusInternalServerError,
				"E9001",
				"Internal server error",
			)

		}
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"S0000",
		"Account closed successfully",
		account,
	)

}
