package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TransactionsHandler) CreateWithdraw(c *gin.Context) {
	accountNumber := c.Param("accountNumber")

	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	transactionEntity := req.ToTransactionEntity()
	transaction, err := h.transactionsService.Withdraw(accountNumber, transactionEntity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Withdraw successful",
		"data":    transaction,
	})
}
