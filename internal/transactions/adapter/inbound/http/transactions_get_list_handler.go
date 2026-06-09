package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TransactionsHandler) GetTransactionList(c *gin.Context) {
	transcations, err := h.transactionsService.GetTransactionList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK,
		ToGetTransactionListResponse(transcations),
	)
}
