package domain

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	ErrorAccountNotFound                          = NewError("Account not found", http.StatusNotFound, nil)
	ErrorTransactionNotFound                      = NewError("Transaction not found", http.StatusNotFound, nil)
	ErrorOperationTypeNotFound                    = NewError("Operation type not found", http.StatusNotFound, nil)
	ErrorAmountMustBePositiveForThatOperationType = NewError("Amount must be positive for that operation type", http.StatusBadRequest, nil)
	ErrorAmountMustBeNegativeForThatOperationType = NewError("Amount must be negative for that operation type", http.StatusBadRequest, nil)
)

func NewError(message string, code int, errors []error) error {
	return &Error{Message: message, Code: code, Errors: errors}
}

type Error struct {
	Message string  `json:"message"`
	Code    int     `json:"-"`
	Errors  []error `json:"errors,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewValidationErrors(ve validator.ValidationErrors) []error {
	out := make([]error, len(ve))
	for i, fe := range ve {
		out[i] = NewValidationError(fe)
	}
	return out
}

func NewValidationError(fieldError validator.FieldError) error {
	return &ValidationError{fieldError.Field(), getErrorMsg(fieldError)}
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "min":
		return "Should be greater or equal than " + fe.Param()
	case "max":
		return "Should be less or equal than " + fe.Param()
	}
	return "Unknown error"
}
