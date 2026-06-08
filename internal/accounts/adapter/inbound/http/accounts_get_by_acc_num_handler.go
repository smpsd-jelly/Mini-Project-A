package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AccountsHandler) GetAccountDetailByAccountNumber(c *gin.Context) {
	accountNumber := c.Param("accountNumber")

	account, err := h.accountsService.GetAccountDetailByAccountNumber(accountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := ToGetAccountDetailResponse(account)

	c.JSON(http.StatusOK, response)

}
