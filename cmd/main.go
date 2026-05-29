package main

import (
	"fmt"
	"os"
	"time"

	accountsHttp "mini-project-a/internal/accounts/adapter/inbound/http"
	accountsPostgres "mini-project-a/internal/accounts/adapter/outbound/postgres"
	accountsCore "mini-project-a/internal/accounts/core"
	"mini-project-a/internal/infrastructure/database"
	transactionsPostgres "mini-project-a/internal/transactions/adapter/outbound/postgres"
	transactionsCore "mini-project-a/internal/transactions/core"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using environment variables")
	}

	db, err := database.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = loc
	router := gin.Default()

	api := router.Group("/api/v1")

	accountsRepo := accountsPostgres.NewAccountPostgresRepository(db)
	transactionsRepo := transactionsPostgres.NewTransactionsPostgresRepository(db)
	transactionsService := transactionsCore.NewTransactionsService(transactionsRepo)

	accountsService := accountsCore.NewAccountsService(
		accountsRepo,
		transactionsService,
	)

	accountsHandler := accountsHttp.NewAccountsHandler(accountsService)

	accountsHttp.RegisterAccountsRoutes(api, accountsHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
