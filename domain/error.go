package domain

import (
	"net/http"
)

var (
	ErrAccountNotFound     = Error("account not found", http.StatusNotFound)
	ErrTransactionNotFound = Error("transaction not found", http.StatusNotFound)
)

func Error(message string, code int) error {
	return &Err{message, code}
}

type Err struct {
	Message string `json:"message"`
	Code    int
}

func (e *Err) Error() string {
	return e.Message
}
