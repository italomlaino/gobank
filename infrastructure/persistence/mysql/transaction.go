package mysql

import (
	"time"

	"github.com/italomlaino/gobank/domain"
)

type TransactionRepository struct {
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(accountID int64, operationTypeID domain.OperationType, amount int64, eventData time.Time) (*domain.Transaction, error) {
	statement, err := DB.Prepare("INSERT INTO transactions(account_id, operation_type_id, amount, event_data) VALUES (?,?,?,?)")
	if err != nil {
		return nil, err
	}

	result, err := statement.Exec(accountID, operationTypeID, amount, eventData)
	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &domain.Transaction{
		ID:              lastId,
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		EventData:       eventData,
	}, nil
}

func (r *TransactionRepository) FetchByAccountID(accountID int64) (*[]domain.Transaction, error) {
	statement, err := DB.Query("SELECT * FROM transactions WHERE account_id = ?", accountID)
	if err != nil {
		return nil, err
	}

	var results []domain.Transaction
	for statement.Next() {
		var id int64
		var accountID int64
		var operationTypeID domain.OperationType
		var amount int64
		var eventData time.Time

		err = statement.Scan(&id, &accountID, &operationTypeID, &amount, &eventData)
		if err != nil {
			return nil, err
		}

		transaction := domain.Transaction{
			ID:              id,
			AccountID:       accountID,
			OperationTypeID: operationTypeID,
			Amount:          amount,
			EventData:       eventData,
		}
		results = append(results, transaction)
	}

	return &results, nil
}
