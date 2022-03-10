package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/italomlaino/gobank/domain"
)

type TransactionController struct {
	*domain.TransactionService
}

type CreateTransactionDTO struct {
	AccountID       int64                `json:"account_id" binding:"required"`
	OperationTypeID domain.OperationType `json:"operation_type_id" binding:"required"`
	Amount          int64                `json:"amount" binding:"required"`
}

func NewTransactionController(service *domain.TransactionService) *TransactionController {
	return &TransactionController{service}
}

func (controller *TransactionController) CreateTransationHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto CreateTransactionDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		transaction, err := controller.TransactionService.Create(dto.AccountID, dto.OperationTypeID, dto.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transaction)
	}
}

func (controller *TransactionController) ListTransationHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		accounts, err := controller.TransactionService.FetchAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, accounts)
	}
}
