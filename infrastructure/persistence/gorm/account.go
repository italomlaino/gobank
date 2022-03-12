package gorm

import (
	"github.com/italomlaino/gobank/domain"
)

type GormAccount struct {
	ID             int64 `gorm:"column:id;primaryKey"`
	DocumentNumber int64 `gorm:"column:document_number"`
}

type GormAccountRepository struct {
}

func NewAccountRepository() *GormAccountRepository {
	return &GormAccountRepository{}
}

func (repository *GormAccountRepository) Create(documentNumber int64) (*domain.Account, error) {
	entity := GormAccount{
		DocumentNumber: documentNumber,
	}
	result := DB.Create(&entity)
	if result.Error != nil {
		return nil, result.Error
	}

	transaction := domain.Account(entity)
	return &transaction, nil
}

func (repository *GormAccountRepository) FetchByID(id int64) (*domain.Account, error) {
	var entity GormAccount
	result := DB.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}

	account := domain.Account(entity)
	return &account, nil
}
