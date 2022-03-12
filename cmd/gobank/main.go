package main

import (
	"os"

	"github.com/italomlaino/gobank/application/controller"
	"github.com/italomlaino/gobank/application/router"
	"github.com/italomlaino/gobank/domain"
	"github.com/italomlaino/gobank/infrastructure/persistence/gorm"
)

func main() {
	gorm.Open()
	defer gorm.Close()

	accountRepository := gorm.NewGormAccountRepository()
	accountService := domain.NewAccountService(accountRepository)
	accountController := controller.NewAccountController(accountService)

	transactionRepository := gorm.NewGormTransactionRepository()
	transactionService := domain.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	port := os.Getenv("PORT")
	router := router.NewRouter(port, accountController, transactionController)
	router.Start()
}
