package gorm

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/italomlaino/gobank/domain"
)

type Transaction struct {
	ID              int64     `gorm:"column:id;primaryKey"`
	AccountID       int64     `gorm:"column:account_id"`
	OperationTypeID int64     `gorm:"column:operation_type_id"`
	Amount          int64     `gorm:"column:amount"`
	EventData       time.Time `gorm:"column:event_data"`
}

type TransactionRepository struct {
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(accountID int64, operationTypeID int64, amount int64, eventData time.Time) (*domain.Transaction, error) {
	var accountExists bool
	err := DB.Model(&Account{}).
		Select("count(*) > 0").
		Where("id = ?", accountID).
		Find(&accountExists).
		Error
	if err != nil {
		return nil, err
	}
	if !accountExists {
		return nil, domain.ErrorAccountNotFound
	}

	var operationTypeExists bool
	err = DB.Model(&OperationType{}).
		Select("count(*) > 0").
		Where("id = ?", operationTypeID).
		Find(&operationTypeExists).
		Error
	if err != nil {
		return nil, err
	}
	if !operationTypeExists {
		return nil, domain.ErrorOperationTypeNotFound
	}

	entity := Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		EventData:       eventData,
	}
	result := DB.Create(&entity)
	if result.Error != nil {
		return nil, result.Error
	}

	transaction := domain.Transaction(entity)
	return &transaction, nil
}

func (r *TransactionRepository) FetchByAccountID(accountID int64) (*[]domain.Transaction, error) {
	var entities []Transaction
	err := DB.Find(&entities, "account_id = ?", accountID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrorTransactionNotFound
		}

		return nil, err
	}

	transactions := make([]domain.Transaction, len(entities))
	for i := range entities {
		transactions[i] = domain.Transaction(entities[i])
	}
	return &transactions, nil
}
