package transaction

import (
	"time"

	"github.com/google/uuid"
)

type TransactionService interface {
	CreateTransaction(userId, categoryId uuid.UUID, amount float64, createdAt time.Time) (*Transaction, error)
	GetTransactionById(transactionId, userId uuid.UUID) (*Transaction, error)
	UpdateTransaction(transactionId, userId, categoryId uuid.UUID, amount *float64) (*Transaction, error)
	DeleteTransaction(transactionId, userId uuid.UUID) error
	GetAllTransactions(userId uuid.UUID) ([]*Transaction, error)
}

type transactionService struct {
	repo TransactionRepository
}

var (
  _ TransactionService = &transactionService{}
)

func NewTransactionService(repo TransactionRepository) *transactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) CreateTransaction(userId, categoryId uuid.UUID, amount float64, createdAt time.Time) (*Transaction, error) {
	if createdAt.String() == "" {
		return nil, ErrDateNull
	}
	if categoryId.String() == "" {
		return nil, ErrTransactionCategoryNull
	}
	if amount == 0.0 {
		return nil, ErrAmountNull
	}
	if amount < 0.0 {
		return nil, ErrAmountNotPositive
	}

	newTransaction := NewTransaction(userId, categoryId, amount, createdAt)
	if newTransaction == nil {
		return nil, ErrTransactionNull
	}
	if err := s.repo.AddTransaction(newTransaction); err != nil {
		return nil, err
	}
	return newTransaction, nil
}

func (s *transactionService) GetTransactionById(transactionId, userId uuid.UUID) (*Transaction, error) {
	if transactionId.String() == "" {
		return nil, ErrTransactionNull
	}
	transaction, err := s.repo.GetTransaction(transactionId, userId)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *transactionService) UpdateTransaction(transactionId, userId, categoryId uuid.UUID, amount *float64) (*Transaction, error) {
	if transactionId.String() == "" || userId.String() == "" || categoryId.String() == "" {
		return nil, ErrTransactionNull
	}
	updateTransaction, err := s.repo.UpdateTransaction(transactionId, userId, categoryId, amount)
	if err != nil {
		return nil, err
	}
	return updateTransaction, nil
}

func (s *transactionService) DeleteTransaction(transactionId, userId uuid.UUID) error {
	if transactionId.String() == "" || userId.String() == "" {
		return ErrTransactionNull
	}

	err := s.repo.DeleteTransaction(transactionId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *transactionService) GetAllTransactions(userId uuid.UUID)([]*Transaction, error ) {
  var transactions []*Transaction
	if userId.String() == "" {
		return transactions, ErrTransactionNull
	}
  transactions, err := s.repo.ListTransactions(userId)
  if err != nil {
    return transactions, err
  }
  return transactions, nil
}
