package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AccountsHandler) CreateAccount(c *gin.Context) {
	var req CreateAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	account := req.ToAccountEntity()

	_, err := h.accountsService.CreateAccount(account)

	if err != nil {
		log.Println("CreateAccount error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "201 Created",
	})
}
