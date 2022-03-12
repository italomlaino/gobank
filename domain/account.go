package domain

type Account struct {
	ID             int64 `json:"id"`
	DocumentNumber int64 `json:"document_number"`
}

type AccountRepository interface {
	Create(documentNumber int64) (*Account, error)
	FetchByID(id int64) (*Account, error)
}

type AccountService interface {
	Create(documentNumber int64) (*Account, error)
	FetchByID(id int64) (*Account, error)
}
