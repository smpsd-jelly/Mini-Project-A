package http

import (
	"errors"
	"mini-project-a/internal/shared/response"
	"mini-project-a/internal/transactions/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TransactionsHandler) CreateDeposit(c *gin.Context) {
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

	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateError(
			c,
			http.StatusBadRequest,
			"E1001",
			"Validation failed",
			response.FormatValidationErrors(err),
		)
		return
	}

	transactionEntity := req.ToTransactionEntity()

	transaction, err := h.transactionsService.Deposit(uri.AccountNumber, transactionEntity)

	if err != nil {
		switch {
		case errors.Is(err, core.ErrAccountNotFound):
			response.Error(
				c,
				http.StatusNotFound,
				"E4001",
				"Account not found",
			)
		case errors.Is(err, core.ErrAccountIsClosed):
			response.Error(
				c,
				http.StatusConflict,
				"E3001",
				"Account is closed",
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
		http.StatusCreated,
		"S0001",
		"Deposit successful",
		transaction,
	)

}
