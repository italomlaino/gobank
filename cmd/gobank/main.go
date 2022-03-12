package main

import (
	"os"

	"github.com/italomlaino/gobank/application/controller"
	"github.com/italomlaino/gobank/application/router"
	"github.com/italomlaino/gobank/domain"
	persistence "github.com/italomlaino/gobank/infrastructure/persistence/gorm"
)

func main() {
	persistence.Open()
	defer persistence.Close()

	accountRepository := persistence.NewAccountRepository()
	accountService := domain.NewAccountService(accountRepository)
	accountController := controller.NewAccountController(accountService)

	transactionRepository := persistence.NewTransactionRepository()
	transactionService := domain.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	port := os.Getenv("PORT")
	router := router.NewRouter(port, accountController, transactionController)
	router.Start()
}
