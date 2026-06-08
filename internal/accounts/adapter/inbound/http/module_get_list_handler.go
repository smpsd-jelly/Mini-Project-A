package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AccountsHandler) GetAccountList(c *gin.Context) {
	accounts, err := h.accountsService.GetAccountList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK,
		ToGetAccountListResponse(accounts),
	)
}
