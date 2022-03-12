package service

import (
	"time"

	"github.com/italomlaino/gobank/domain"
)

type TransactionService struct {
	domain.TransactionRepository
}

func NewTransactionService(repository domain.TransactionRepository) *TransactionService {
	return &TransactionService{repository}
}

func (s *TransactionService) Create(accountID int64, operationTypeID int64, amount int64) (*domain.Transaction, error) {
	return s.TransactionRepository.Create(accountID, operationTypeID, amount, time.Now())
}

func (s *TransactionService) FetchByAccountID(accountID int64) (*[]domain.Transaction, error) {
	return s.TransactionRepository.FetchByAccountID(accountID)
}
