package main

import (
	"os"

	"github.com/italomlaino/gobank/controller"
	"github.com/italomlaino/gobank/domain"
	"github.com/italomlaino/gobank/infrastructure/persistence/mysql"
	"github.com/italomlaino/gobank/infrastructure/server"
)

func main() {
	accountRepository := mysql.NewMysqlAccountRepository()
	accountService := domain.NewAccountService(accountRepository)
	accountController := controller.NewAccountController(accountService)

	transactionRepository := mysql.NewMysqlTransactionRepository()
	transactionService := domain.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	port := os.Getenv("PORT")
	server := server.NewServer(port, accountController, transactionController)
	server.Start()
}
