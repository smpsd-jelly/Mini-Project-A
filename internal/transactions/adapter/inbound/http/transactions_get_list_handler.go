package http

import (
	"errors"
	"mini-project-a/internal/shared/response"
	"mini-project-a/internal/transactions/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TransactionsHandler) GetTransactionList(c *gin.Context) {
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

	transactions, err := h.transactionsService.GetTransactionList(uri.AccountNumber)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			response.Error(
				c,
				http.StatusNotFound,
				"E4001",
				"Account not found",
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
		"S0001",
		"Transactions retrieved successfully",
		transactions,
	)
}
