package http

import "github.com/gin-gonic/gin"

func RegisterAccountsRoutes(router *gin.RouterGroup, accountsHandler *AccountsHandler) {
	accounts := router.Group("/accounts")

	accounts.POST("", accountsHandler.CreateAccount)
	accounts.GET("/:accountNumber", accountsHandler.GetAccountDetailByAccountNumber)
	accounts.GET("", accountsHandler.GetAccountList)
	accounts.PATCH("/:accountNumber/close", accountsHandler.CloseAccount)
	accounts.PATCH("/reopen/:accountNumber", accountsHandler.ReopenAccountHandler)
}
