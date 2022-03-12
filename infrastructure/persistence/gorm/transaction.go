package gorm

import (
	"time"

	"github.com/italomlaino/gobank/domain"
)

type Transaction struct {
	ID              int64                `gorm:"column:id;primaryKey"`
	AccountID       int64                `gorm:"column:account_id"`
	OperationTypeID domain.OperationType `gorm:"column:operation_type_id"`
	Amount          int64                `gorm:"column:amount"`
	EventData       time.Time            `gorm:"column:event_data"`
}

type TransactionRepository struct {
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(accountID int64, operationTypeID domain.OperationType, amount int64, eventData time.Time) (*domain.Transaction, error) {
	var exists bool
	err := DB.Model(&Account{}).
		Select("count(*) > 0").
		Where("id = ?", accountID).
		Find(&exists).
		Error
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, domain.ErrAccountNotFound
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
	result := DB.Find(&entities, "account_id = ?", accountID)
	if result.Error != nil {
		return nil, result.Error
	}

	transactions := make([]domain.Transaction, len(entities))
	for i := range entities {
		transactions[i] = domain.Transaction(entities[i])
	}
	return &transactions, nil
}
