package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

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
			var parsedError *domain.Error
			switch err.(type) {
			case *domain.Error:
				parsedError = err.(*domain.Error)
			case validator.ValidationErrors:
				var ve validator.ValidationErrors
				errors.As(err, &ve)
				validationErrors := domain.NewValidationErrors(ve)
				parsedError = domain.NewError("Field Validation Error", http.StatusBadRequest, validationErrors).(*domain.Error)
			default:
				parsedError = domain.NewError("Internal Server Error", http.StatusInternalServerError, nil).(*domain.Error)
			}
			c.IndentedJSON(parsedError.Code, parsedError)
			c.Abort()
			return
		}

	}
}
