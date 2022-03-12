package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italomlaino/gobank/domain"
)

func ErrorMiddleware() gin.HandlerFunc {
	return createErrorHandler(gin.ErrorTypeAny)
}

func createErrorHandler(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		detectedErrors := c.Errors.ByType(errType)
		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *domain.Err
			switch err.(type) {
			case *domain.Err:
				parsedError = err.(*domain.Err)
			default:
				parsedError = domain.Error("Internal Server Error", http.StatusInternalServerError).(*domain.Err)
			}
			c.IndentedJSON(parsedError.Code, parsedError)
			c.Abort()
			return
		}

	}
}
