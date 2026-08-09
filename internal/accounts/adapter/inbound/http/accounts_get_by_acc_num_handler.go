package http

import (
	"errors"
	"mini-project-a/internal/accounts/core"
	responses "mini-project-a/internal/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AccountsHandler) GetAccountDetailByAccountNumber(c *gin.Context) {
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

	account, err := h.accountsService.GetAccountDetailByAccountNumber(uri.AccountNumber)

	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			responses.Error(
				c,
				http.StatusNotFound,
				"E4001",
				"Account not found",
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

	response := ToGetAccountDetailResponse(account)

	responses.Success(
		c,
		http.StatusOK,
		"S0000",
		"Account retrieved successfully",
		response,
	)

}
