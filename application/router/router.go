package router

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/italomlaino/gobank/application/controller"
)

type Router struct {
	controller.AccountController
	controller.TransactionController

	port string
}

func NewRouter(port string, accountController controller.AccountController, transactionController controller.TransactionController) *Router {
	return &Router{accountController, transactionController, port}
}

func (r *Router) Start() {
	log.Printf("Listening on port %s", r.port)
	log.Printf("Open http://localhost:%s in the browser", r.port)

	router := r.create()
	log.Fatal(router.Run(":" + r.port))
}

func (r *Router) create() *gin.Engine {
	router := gin.Default()
	router.Use(ErrorMiddleware())
	router.POST("/accounts", r.AccountController.CreateAccountHandler())
	router.GET("/accounts/:accountId", r.AccountController.FetchAccountHandler())
	router.POST("/transactions", r.TransactionController.CreateTransationHandler())
	router.GET("/transactions", r.TransactionController.ListTransationHandler())
	return router
}
