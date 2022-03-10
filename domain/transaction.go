package domain

import "time"

type OperationType int64

const (
	Purchase OperationType = iota
	PurchaseInInstallments
	Withdraw
	Payment
)

type Transaction struct {
	ID              int64         `json:"id"`
	AccountID       int64         `json:"account_id"`
	OperationTypeID OperationType `json:"operation_type_id"`
	Amount          int64         `json:"amount"`
	EventData       time.Time     `json:"event_data"`
}

type TransactionRepository interface {
	Create(accountID int64, operationTypeID OperationType, amount int64, eventData time.Time) (*Transaction, error)
	FetchAll() (*[]Transaction, error)
}

type TransactionService struct {
	repository TransactionRepository
}

func NewTransactionService(repository TransactionRepository) TransactionService {
	return TransactionService{repository}
}

func (service TransactionService) Create(accountID int64, operationTypeID OperationType, amount int64) (*Transaction, error) {
	return service.repository.Create(accountID, operationTypeID, amount, time.Now())
}

func (service TransactionService) FetchAll() (*[]Transaction, error) {
	return service.repository.FetchAll()
}
