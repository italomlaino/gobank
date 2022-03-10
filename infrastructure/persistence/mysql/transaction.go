package mysql

import (
	"time"

	"github.com/italomlaino/gobank/domain"
)

type MysqlTransactionRepository struct {
}

func NewMysqlTransactionRepository() *MysqlTransactionRepository {
	return &MysqlTransactionRepository{}
}

func (repository *MysqlTransactionRepository) Create(accountID int64, operationTypeID domain.OperationType, amount int64, eventData time.Time) (*domain.Transaction, error) {
	db := Connect()
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO account_transaction(account_id, operation_type_id, amount, event_data) VALUES (?,?,?,?)")
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

func (repository *MysqlTransactionRepository) FetchAll() (*[]domain.Transaction, error) {
	db := Connect()
	defer db.Close()

	statement, err := db.Query("SELECT * FROM account_transaction")
	if err != nil {
		return nil, err
	}

	results := []domain.Transaction{}
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
