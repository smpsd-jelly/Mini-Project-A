package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AccountsHandler) CloseAccount(c *gin.Context) {
	accountNumber := c.Param("accountNumber")

	_, err := h.accountsService.CloseAccount(accountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "200 OK",
	})
}
