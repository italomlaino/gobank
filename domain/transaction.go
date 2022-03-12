package domain

import "time"

type Transaction struct {
	ID              int64           `json:"id"`
	AccountID       int64           `json:"account_id"`
	OperationTypeID OperationTypeID `json:"operation_type_id"`
	Amount          int64           `json:"amount"`
	EventData       time.Time       `json:"event_data"`
}

type TransactionRepository interface {
	Create(accountID int64, operationTypeID OperationTypeID, amount int64, eventData time.Time) (*Transaction, error)
	FetchByAccountID(accountID int64) (*[]Transaction, error)
}

type TransactionService interface {
	Create(accountID int64, operationTypeID OperationTypeID, amount int64) (*Transaction, error)
	FetchByAccountID(accountID int64) (*[]Transaction, error)
}
