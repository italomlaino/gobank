package mysql

import (
	"github.com/italomlaino/gobank/domain"
)

type MysqlAccountRepository struct {
}

func NewMysqlAccountRepository() *MysqlAccountRepository {
	return &MysqlAccountRepository{}
}

func (repository *MysqlAccountRepository) Create(documentNumber int64) (*domain.Account, error) {
	db := Connect()
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO accounts(document_number) VALUES (?)")
	if err != nil {
		return nil, err
	}

	result, err := statement.Exec(documentNumber)
	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &domain.Account{
		ID:             lastId,
		DocumentNumber: documentNumber,
	}, nil
}

func (repository *MysqlAccountRepository) FetchByID(id int64) (*domain.Account, error) {
	db := Connect()
	defer db.Close()

	statement, err := db.Query("SELECT * FROM accounts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	if statement.Next() {
		var id int64
		var documentNumber int64

		err = statement.Scan(&id, &documentNumber)
		if err != nil {
			return nil, err
		}

		account := domain.Account{
			ID:             id,
			DocumentNumber: documentNumber,
		}
		return &account, nil
	}

	return nil, nil
}
