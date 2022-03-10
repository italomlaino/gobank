package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/italomlaino/gobank/domain"
)

type CreateAccountDTO struct {
	DocumentNumber int64 `json:"document_number" binding:"required"`
}

type FetchAccountByIdDTO struct {
	ID int64 `uri:"accountId" binding:"required"`
}

type AccountController interface {
	CreateAccountHandler() func(c *gin.Context)
	FetchAccountHandler() func(c *gin.Context)
}

type DefaultAccountController struct {
	domain.AccountService
}

func NewAccountController(service domain.AccountService) *DefaultAccountController {
	return &DefaultAccountController{service}
}

func (controller *DefaultAccountController) CreateAccountHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto CreateAccountDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		account, err := controller.AccountService.Create(dto.DocumentNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
	}
}

func (controller *DefaultAccountController) FetchAccountHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto FetchAccountByIdDTO
		if err := c.ShouldBindUri(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		account, err := controller.AccountService.FetchByID(dto.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
	}
}
