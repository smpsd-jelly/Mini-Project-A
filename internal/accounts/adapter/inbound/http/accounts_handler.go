package http

import (
	"database/sql"
	"net/http"

	"mini-project-a/internal/accounts/core"

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

	account := req.ToAccountEntity()

	_, err := h.accountsService.CreateAccount(account)
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

func (h *AccountsHandler) GetAccountDetailByAccountNumber(c *gin.Context) {
	accountNumber := c.Param("accountNumber")

	account, err := h.accountsService.GetAccountDetailByAccountNumber(accountNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Account not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	response := ToGetAccountDetailResponse(account)

	c.JSON(http.StatusOK, response)
}
