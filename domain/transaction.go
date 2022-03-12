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
	FetchByAccountID(accountID int64) (*[]Transaction, error)
}

type TransactionService interface {
	Create(accountID int64, operationTypeID OperationType, amount int64) (*Transaction, error)
	FetchByAccountID(accountID int64) (*[]Transaction, error)
}

type DefaultTransactionService struct {
	TransactionRepository
}

func NewTransactionService(repository TransactionRepository) *DefaultTransactionService {
	return &DefaultTransactionService{repository}
}

func (service *DefaultTransactionService) Create(accountID int64, operationTypeID OperationType, amount int64) (*Transaction, error) {
	return service.TransactionRepository.Create(accountID, operationTypeID, amount, time.Now())
}

func (service *DefaultTransactionService) FetchByAccountID(accountID int64) (*[]Transaction, error) {
	return service.TransactionRepository.FetchByAccountID(accountID)
}
