package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TransactionsHandler) GetTransactionList(c *gin.Context) {
	accountNumber := c.Param("accountNumber")

	transactions, err := h.transactionsService.GetTransactionList(accountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK,
		ToGetTransactionListResponse(transactions),
	)
}
