package domain

type Account struct {
	ID             int64 `json:"id"`
	DocumentNumber int64 `json:"document_number"`
}

type AccountRepository interface {
	Create(documentNumber int64) (*Account, error)
	Delete(id int64) error
	FetchByDocumentNumber(documentNumber int64) (*Account, error)
	FetchByID(id int64) (*Account, error)
}

type AccountService struct {
	repository AccountRepository
}

func NewAccountService(repository AccountRepository) AccountService {
	return AccountService{repository}
}

func (service AccountService) Create(documentNumber int64) (*Account, error) {
	return service.repository.Create(documentNumber)
}

func (service AccountService) Delete(id int64) error {
	service.repository.Delete(id)
	return nil
}

func (service AccountService) FetchByDocumentNumber(documentNumber int64) (*Account, error) {
	return service.repository.FetchByDocumentNumber(documentNumber)
}

func (service AccountService) FetchByID(id int64) (*Account, error) {
	return service.repository.FetchByID(id)
}
