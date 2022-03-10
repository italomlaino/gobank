package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/italomlaino/gobank/controller"
	"github.com/italomlaino/gobank/domain"
	"github.com/italomlaino/gobank/infrastructure/persistence/mysql"
)

func main() {
	startServer()
}

func startServer() {
	port := os.Getenv("PORT")
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)

	router := createRouter()
	log.Fatal(router.Run(":" + port))
}

func createRouter() *gin.Engine {
	accountRepository := mysql.NewMysqlAccountRepository()
	accountService := domain.NewAccountService(accountRepository)
	accountController := controller.NewAccountController(accountService)

	transactionRepository := mysql.NewMysqlTransactionRepository()
	transactionService := domain.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	router := gin.Default()
	router.POST("/accounts", accountController.CreateAccountHandler())
	router.GET("/accounts/:accountId", accountController.FetchAccountHandler())
	router.POST("/transactions", transactionController.CreateTransationHandler())
	router.GET("/transactions", transactionController.ListTransationHandler())
	return router
}
