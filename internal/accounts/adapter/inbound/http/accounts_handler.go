package http

import (
	"net/http"
	"time"

	"mini-project-a/internal/accounts/core"
	"mini-project-a/internal/accounts/core/entity"

	"github.com/gin-gonic/gin"
)

type AccountsHandler struct {
	accountsService *core.AccountsService
}

func NewAccountsHandler(accountsService *core.AccountsService) *AccountsHandler {
	return &AccountsHandler{
		accountsService: accountsService,
	}
}

func (h *AccountsHandler) CreateAccount(c *gin.Context) {
	var req CreateAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	account := &entity.Accounts{
		OwnerName:   req.OwnerName,
		CitizenID:   req.CitizenID,
		PhoneNumber: req.PhoneNumber,
		AccountType: req.AccountType,
		Balance:     req.Balance,
		Status:      "ACTIVE",
		CreatedAt:   time.Now(),
	}

	err := h.accountsService.CreateAccount(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "201 Created",
	})
}
