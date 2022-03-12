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

func NewGormAccountRepository() *GormAccountRepository {
	return &GormAccountRepository{}
}

func (repository *GormAccountRepository) Create(documentNumber int64) (*domain.Account, error) {
	db, sqlDB, err := Open()
	if err != nil {
		return nil, err
	}
	defer sqlDB.Close()

	entity := GormAccount{
		DocumentNumber: documentNumber,
	}
	result := db.Create(&entity)
	if result.Error != nil {
		return nil, result.Error
	}

	transaction := domain.Account(entity)
	return &transaction, nil
}

func (repository *GormAccountRepository) FetchByID(id int64) (*domain.Account, error) {
	db, sqlDB, err := Open()
	if err != nil {
		return nil, err
	}
	defer sqlDB.Close()

	var entity GormAccount
	result := db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}

	account := domain.Account(entity)
	return &account, nil
}
