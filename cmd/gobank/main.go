package main

import (
	"os"

	"github.com/italomlaino/gobank/application/controller"
	"github.com/italomlaino/gobank/application/router"
	"github.com/italomlaino/gobank/domain"
	"github.com/italomlaino/gobank/infrastructure/persistence/mysql"
)

func main() {
	accountRepository := mysql.NewMysqlAccountRepository()
	accountService := domain.NewAccountService(accountRepository)
	accountController := controller.NewAccountController(accountService)

	transactionRepository := mysql.NewMysqlTransactionRepository()
	transactionService := domain.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	port := os.Getenv("PORT")
	server := router.NewRouter(port, accountController, transactionController)
	server.Start()
}
