package http

import "github.com/gin-gonic/gin"

func TransactionsRoutes(router *gin.RouterGroup, transactionsHandler *TransactionsHandler) {
	accounts := router.Group("/accounts")
	accounts.POST("/:accountNumber/deposit", transactionsHandler.CreateDeposit)
	accounts.POST("/:accountNumber/withdraw", transactionsHandler.CreateWithdraw)

}
