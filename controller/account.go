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

type AccountController struct {
	service domain.AccountService
}

func NewAccountController(service domain.AccountService) AccountController {
	return AccountController{service}
}

func (controller AccountController) CreateAccountHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto CreateAccountDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		account, err := controller.service.Create(dto.DocumentNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
	}
}

func (controller AccountController) FetchAccountHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto FetchAccountByIdDTO
		if err := c.ShouldBindUri(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		account, err := controller.service.FetchByID(dto.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
	}
}
