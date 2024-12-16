package transaction

import (
	"fmt"
	"time"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	AddTransaction(transaction *Transaction) error
	GetTransaction(transactionId, userId uuid.UUID) (*Transaction, error)
	UpdateTransaction(transactionId, userId, categoryId uuid.UUID, amount float64) (*Transaction, error)
	DeleteTransaction(transactionId, userId uuid.UUID) error
	ListTransactions(userId uuid.UUID) ([]*Transaction, error)
}

type InMemoryStore struct {
	Store map[uuid.UUID]map[uuid.UUID]*Transaction
}

type Transaction struct {
	Id         uuid.UUID
	UserID     uuid.UUID
	CategoryID uuid.UUID
	Amount     float64
	CreatedAt  time.Time
	updatedAt  time.Time
}

var (
	_ TransactionRepository = &InMemoryStore{}
)

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{Store: make(map[uuid.UUID]map[uuid.UUID]*Transaction)}
}

func NewTransaction(categoryId, userId uuid.UUID, amount float64, createdAt time.Time) (*Transaction, error) {
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
  t := &Transaction{
  Id: utils.GenerateUUID(),
    CreatedAt: createdAt,
    UserID: userId,
    CategoryID: categoryId,
    Amount: amount,
    updatedAt: time.Now()}
		fmt.Println("Creating transaction with Id: ", t.Id)
	return t, nil
}

func (s *InMemoryStore) AddTransaction(transaction *Transaction) error {
	if transaction == nil {
		return ErrTransactionNull
	}
  userTransactions, ok := s.Store[transaction.UserID]
  if !ok {
    userTransactions = make(map[uuid.UUID]*Transaction)
    s.Store[transaction.UserID] = userTransactions
  }
  userTransactions[transaction.Id] = transaction
	return nil
}

func (s *InMemoryStore) GetTransaction(transactionId, userId uuid.UUID) (*Transaction, error) {
	userTransactions, ok := s.Store[userId]
	if !ok {
		return nil, ErrTransactionWithUserNotFound
	}

  transaction, ok := userTransactions[transactionId]
  if !ok {
    return nil, ErrTransactionNotFound
  }

  return transaction, nil
}


func (s *InMemoryStore) DeleteTransaction(transactionId, userId uuid.UUID) error {
	userTransactions, ok := s.Store[userId]
	if !ok {
		return ErrTransactionWithUserNotFound
	}
	delete(userTransactions, transactionId)
	fmt.Printf("Transaction with Id '%s' has been deleted", transactionId)

	return nil
}

func (s *InMemoryStore) UpdateTransaction(transactionId, userId, categoryId uuid.UUID, amount float64) (*Transaction, error) {
	userTransactions, ok := s.Store[userId]
	if !ok {
		return nil, ErrTransactionWithUserNotFound
	}
  transaction, ok := userTransactions[transactionId]
  if !ok {
    return nil, ErrTransactionNotFound
  }
	if transaction.CategoryID.String() == "" {
		return nil, ErrTransactionCategoryNull
	}
	if transaction.CategoryID != categoryId {
		transaction.CategoryID = categoryId
	}
	if transaction.Amount != amount {
		transaction.Amount = amount
	}
	transaction.updatedAt = time.Now()

	return transaction, nil
}

func (s *InMemoryStore) ListTransactions(userId uuid.UUID) ([]*Transaction, error) {
  var result []*Transaction
	fmt.Println("Getting transactions...")
  userTransactions, ok := s.Store[userId]
  if !ok {
    return result, ErrTransactionWithUserNotFound
  }
	for _, transaction := range userTransactions {
		result = append(result, transaction)
	}
	return result, nil
}
