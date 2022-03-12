package router

import (
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

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
	r.setupBinding()
	log.Fatal(router.Run(":" + r.port))
}

func (r *Router) create() *gin.Engine {
	router := gin.Default()
	router.Use(ErrorMiddleware())
	router.POST("/accounts", r.AccountController.CreateHandler())
	router.GET("/accounts/:accountId", r.AccountController.FetchByIDHandler())
	router.GET("/accounts/:accountId/transactions", r.TransactionController.FetchByAccountIDHandler())
	router.POST("/transactions", r.TransactionController.CreateHandler())
	return router
}

func (r *Router) setupBinding() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
