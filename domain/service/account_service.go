package service

import "github.com/italomlaino/gobank/domain"

type AccountService struct {
	domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository}
}

func (s *AccountService) Create(documentNumber int64) (*domain.Account, error) {
	return s.AccountRepository.Create(documentNumber)
}

func (s *AccountService) FetchByID(id int64) (*domain.Account, error) {
	return s.AccountRepository.FetchByID(id)
}
