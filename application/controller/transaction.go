package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/italomlaino/gobank/domain"
)

type CreateTransactionDTO struct {
	AccountID       int64                `json:"account_id" binding:"required,number,min=1"`
	OperationTypeID domain.OperationType `json:"operation_type_id" binding:"required,number,min=1,max=4"`
	Amount          int64                `json:"amount" binding:"required,number,min=1"`
}

type FetchTransactionByAccountIdDTO struct {
	AccountID int64 `uri:"accountId" binding:"required,number,min=1"`
}

type TransactionController interface {
	CreateHandler() func(c *gin.Context)
	FetchByAccountIDHandler() func(c *gin.Context)
}

type DefaultTransactionController struct {
	domain.TransactionService
}

func NewTransactionController(service domain.TransactionService) *DefaultTransactionController {
	return &DefaultTransactionController{service}
}

func (con *DefaultTransactionController) CreateHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto CreateTransactionDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.Error(err)
			return
		}

		transaction, err := con.TransactionService.Create(dto.AccountID, dto.OperationTypeID, dto.Amount)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, transaction)
	}
}

func (con *DefaultTransactionController) FetchByAccountIDHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto FetchTransactionByAccountIdDTO
		if err := c.ShouldBindUri(&dto); err != nil {
			c.Error(err)
			return
		}

		accounts, err := con.TransactionService.FetchByAccountID(dto.AccountID)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, accounts)
	}
}
