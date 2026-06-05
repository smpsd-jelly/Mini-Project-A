package http

import "github.com/gin-gonic/gin"

func RegisterAccountsRoutes(router *gin.RouterGroup, accountsHandler *AccountsHandler) {
	accounts := router.Group("/accounts")

	accounts.POST("", accountsHandler.CreateAccount)
	accounts.GET("/:accountNumber", accountsHandler.GetAccountDetailByAccountNumber)
}
