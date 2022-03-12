package gorm

import (
	"errors"

	"gorm.io/gorm"

	"github.com/italomlaino/gobank/domain"
)

type Account struct {
	ID             int64 `gorm:"column:id;primaryKey"`
	DocumentNumber int64 `gorm:"column:document_number"`
}

type AccountRepository struct {
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) Create(documentNumber int64) (*domain.Account, error) {
	entity := Account{
		DocumentNumber: documentNumber,
	}
	result := DB.Create(&entity)
	if result.Error != nil {
		return nil, result.Error
	}

	transaction := domain.Account(entity)
	return &transaction, nil
}

func (r *AccountRepository) FetchByID(id int64) (*domain.Account, error) {
	var entity Account
	err := DB.First(&entity, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrorAccountNotFound
		}

		return nil, err
	}

	account := domain.Account(entity)
	return &account, nil
}
