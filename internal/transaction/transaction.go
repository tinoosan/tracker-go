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
	GetTransactionByUUID(u uuid.UUID) (*Transaction, error)
	ListTransactions() []Transaction
}

type TransactionsMap struct {
	Store map[uuid.UUID]*Transaction
}

type Transaction struct {
	Id        uuid.UUID
	CreatedAt time.Time
	Category  *category.Category
	Amount    float64
	created   bool
}

var (
	_ TransactionRepository = &TransactionsMap{}
)

func NewTransactionsMap() *TransactionsMap {
	return &TransactionsMap{Store: make(map[uuid.UUID]*Transaction)}
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
	t.created = true
	fmt.Println("Creating transaction with Id: ", t.Id)
	return t, nil
}

func (tm *TransactionsMap) AddTransaction(t Transaction) error {
	if t == (Transaction{}) {
		return ErrTransactionNull
	}
	for {
		_, ok := tm.Store[t.Id]
		if !ok {
			break
		}
		t.Id = utils.GenerateUUID()
	}

	tm.Store[t.Id] = &t
	return nil
}

func (tm *TransactionsMap) GetTransactionByUUID(u uuid.UUID) (*Transaction, error) {
	_, ok := tm.Store[u]
	if ok {
		return tm.Store[u], nil
	}
	return nil, ErrTransactionNotFound
}

func (tm *TransactionsMap) ListTransactions() []Transaction {
	var result []Transaction
	fmt.Println("Getting transactions...")
	for _, v := range tm.Store {
		result = append(result, *v)
	}
	return result
}
