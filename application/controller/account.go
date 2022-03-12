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
	CreateHandler() func(c *gin.Context)
	FetchByIDHandler() func(c *gin.Context)
}

type DefaultAccountController struct {
	domain.AccountService
}

func NewAccountController(service domain.AccountService) *DefaultAccountController {
	return &DefaultAccountController{service}
}

func (con *DefaultAccountController) CreateHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto CreateAccountDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		account, err := con.AccountService.Create(dto.DocumentNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
	}
}

func (con *DefaultAccountController) FetchByIDHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto FetchAccountByIdDTO
		if err := c.ShouldBindUri(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		account, err := con.AccountService.FetchByID(dto.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if account == nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, account)
	}
}
