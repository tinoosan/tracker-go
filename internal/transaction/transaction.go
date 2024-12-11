package transaction

import (
	"fmt"
	"time"
	"trackergo/internal/category"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	AddTransaction(t Transaction) error
	GetTransaction(u uuid.UUID) (*Transaction, error)
	UpdateTransaction(u uuid.UUID, category category.Category, amount float64) (*Transaction, error)
	DeleteTransaction(u uuid.UUID) error
	ListTransactions() []Transaction
}

type InMemoryStore struct {
	Store map[uuid.UUID]*Transaction
}

type Transaction struct {
	Id        uuid.UUID
	CreatedAt time.Time
	Category  *category.Category
	Amount    float64
  updatedAt time.Time
}

var (
	_ TransactionRepository = &InMemoryStore{}
)

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{Store: make(map[uuid.UUID]*Transaction)}
}

func (t *Transaction) NewTransaction(createdAt time.Time, category *category.Category, amount float64) (*Transaction, error) {
	if createdAt.String() == "" {
		return nil, ErrDateNull
	}
	if category == nil {
		return nil, ErrTransactionCategoryNull
	}
	if amount == 0.0 {
		return nil, ErrAmountNull
	}
	if amount < 0.0 {
		return nil, ErrAmountNotPositive
	}
	t.Id = utils.GenerateUUID()
	t.CreatedAt = createdAt
	t.Category = category
	t.Amount = amount
  t.updatedAt = time.Now()
	fmt.Println("Creating transaction with Id: ", t.Id)
	return t, nil
}

func (s *InMemoryStore) AddTransaction(t Transaction) error {
	if t == (Transaction{}) {
		return ErrTransactionNull
	}
	for {
		_, ok := s.Store[t.Id]
		if !ok {
			break
		}
		t.Id = utils.GenerateUUID()
	}

	s.Store[t.Id] = &t
	return nil
}

func (s *InMemoryStore) GetTransaction(u uuid.UUID) (*Transaction, error) {
	_, ok := s.Store[u]
	if ok {
		return s.Store[u], nil
	}
	return nil, ErrTransactionNotFound
}

func (s *InMemoryStore) DeleteTransaction(u uuid.UUID) error {
	val, ok := s.Store[u]
	if !ok {
		return ErrTransactionNotFound
	}
	delete(s.Store, u)
	fmt.Printf("Transaction with Id '%s' has been deleted", val.Id)

	return nil
}

func (s *InMemoryStore) UpdateTransaction(u uuid.UUID, category category.Category, amount float64) (*Transaction, error) {
	val, ok := s.Store[u]
	if !ok {
		return nil, ErrTransactionNotFound
	}
  if val.Category == nil {
    return nil, ErrTransactionCategoryNull
  }
  if val.Category.Name != category.Name {
    val.Category.Name = category.Name
  }
  if val.Amount != amount {
    val.Amount = amount
  }
  val.updatedAt = time.Now()

  return val, nil
}

func (s *InMemoryStore) ListTransactions() []Transaction {
	var result []Transaction
	fmt.Println("Getting transactions...")
	for _, v := range s.Store {
		result = append(result, *v)
	}
	return result
}
