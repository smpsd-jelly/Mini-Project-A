package http

import "github.com/gin-gonic/gin"

func TransactionsRoutes(router *gin.RouterGroup, transactionsHandler *TransactionsHandler) {
	transactions := router.Group("/accounts")
	transactions.POST("/:accountNumber/deposit", transactionsHandler.CreateDeposit)
}
