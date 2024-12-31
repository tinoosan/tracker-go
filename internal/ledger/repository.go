package ledger

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	_ LedgerRepository = &InMemoryStore{}
)

type LedgerRepository interface {
	Save(transaction *Entry) error
	FindByID(transactionId, userId uuid.UUID) (*Entry, error)
	Update(transactionId, userId uuid.UUID, amount *float64) error
	Delete(transactionId, userId uuid.UUID) error
	List(userId uuid.UUID) ([]*Entry, error)
}

type InMemoryStore struct {
	Store map[uuid.UUID]map[uuid.UUID]*Entry
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{Store: make(map[uuid.UUID]map[uuid.UUID]*Entry)}
}

func (s *InMemoryStore) Save(transaction *Entry) error {

	userEntries, ok := s.Store[transaction.UserID]
	if !ok {
		userEntries = make(map[uuid.UUID]*Entry)
		s.Store[transaction.UserID] = userEntries
	}
	userEntries[transaction.ID] = transaction
	fmt.Printf("Entry with ID '%s' has been added to InMemoryStore\n", transaction.ID)
	return nil
}

func (s *InMemoryStore) FindByID(transactionId, userId uuid.UUID) (*Entry, error) {
	userEntries, ok := s.Store[userId]
	if !ok {
		return nil, errors.New("no entries found")
	}

	transaction, ok := userEntries[transactionId]
	if !ok {
		return nil, ErrEntryNotFound
	}

	return transaction, nil
}

func (s *InMemoryStore) Delete(transactionId, userId uuid.UUID) error {
	userEntries, ok := s.Store[userId]
	if !ok {
		return errors.New("no entries found")	
  }

	txn, ok := userEntries[transactionId]
	if !ok {
		return ErrEntryNotFound
	}
	reversedtnx := txn.Reverse()
	s.Save(reversedtnx)

	fmt.Printf("Entry with Id '%s' has been reversed", transactionId)

	return nil
}

func (s *InMemoryStore) Update(transactionId, userId uuid.UUID, amount *float64) error {
	userEntries, ok := s.Store[userId]
	if !ok {
		return errors.New("no entries found")
	}
	transaction, ok := userEntries[transactionId]
	if !ok {
		return ErrEntryNotFound
	}

	if transaction.Amount == *amount && amount != nil {
		reversedTxn, updatedTxn := transaction.UpdateAmount(*amount)
		s.Save(reversedTxn)
		s.Save(updatedTxn)
		return nil
	}

	return nil
}

func (s *InMemoryStore) List(userId uuid.UUID) ([]*Entry, error) {
	var result []*Entry
	fmt.Println("Getting transactions...")
	userEntries, ok := s.Store[userId]
	if !ok {
		return result, errors.New("no entries found")
	}
	for _, transaction := range userEntries {
		result = append(result, transaction)
	}
	return result, nil
}
