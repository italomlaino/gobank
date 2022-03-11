package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/italomlaino/gobank/application/controller"
)

type Server struct {
	controller.AccountController
	controller.TransactionController

	port string
}

func NewServer(port string, accountController controller.AccountController, transactionController controller.TransactionController) *Server {
	return &Server{accountController, transactionController, port}
}

func (server *Server) Start() {
	log.Printf("Listening on port %s", server.port)
	log.Printf("Open http://localhost:%s in the browser", server.port)

	router := server.createRouter()
	log.Fatal(router.Run(":" + server.port))
}

func (server *Server) createRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/accounts", server.AccountController.CreateAccountHandler())
	router.GET("/accounts/:accountId", server.AccountController.FetchAccountHandler())
	router.POST("/transactions", server.TransactionController.CreateTransationHandler())
	router.GET("/transactions", server.TransactionController.ListTransationHandler())
	return router
}
