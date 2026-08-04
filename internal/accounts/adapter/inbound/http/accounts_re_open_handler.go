package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AccountsHandler) ReopenAccountHandler(c *gin.Context) {
	acc := c.Param("accountNumber")
	accounts, err := h.accountsService.ReopenAccount(acc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := ToReOpenAccountDetailResponse(accounts)
	c.JSON(http.StatusOK, response)

}
