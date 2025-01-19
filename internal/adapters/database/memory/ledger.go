package memory

import (
	"errors"
	"fmt"
	"trackergo/internal/domain/ledger"

	"github.com/google/uuid"
)

type LedgerMemoryStore struct {
	Store map[uuid.UUID]map[uuid.UUID]*ledger.Entry
}

func NewLedgerMemoryStore() *LedgerMemoryStore {
	return &LedgerMemoryStore{Store: make(map[uuid.UUID]map[uuid.UUID]*ledger.Entry)}
}

func (s *LedgerMemoryStore) Save(transaction *ledger.Entry) error {

	userEntries, ok := s.Store[transaction.UserID]
	if !ok {
		userEntries = make(map[uuid.UUID]*ledger.Entry)
		s.Store[transaction.UserID] = userEntries
	}
	userEntries[transaction.ID] = transaction
	fmt.Printf("Entry with ID '%s' has been added to LedgerMemoryStore\n", transaction.ID)
	return nil
}

func (s *LedgerMemoryStore) FindByID(transactionId, userId uuid.UUID) (*ledger.Entry, error) {
	userEntries, ok := s.Store[userId]
	if !ok {
		return nil, errors.New("no entries found")
	}

	transaction, ok := userEntries[transactionId]
	if !ok {
		return nil, ledger.ErrEntryNotFound
	}

	return transaction, nil
}

func (s *LedgerMemoryStore) Delete(transactionId, userId uuid.UUID) error {
	userEntries, ok := s.Store[userId]
	if !ok {
		return errors.New("no entries found")
	}

	txn, ok := userEntries[transactionId]
	if !ok {
		return ledger.ErrEntryNotFound
	}
	reversedtnx := txn.Reverse()
	s.Save(reversedtnx)

	fmt.Printf("Entry with Id '%s' has been reversed", transactionId)

	return nil
}

func (s *LedgerMemoryStore) Update(transactionId, userId uuid.UUID, amount *float64) error {
	userEntries, ok := s.Store[userId]
	if !ok {
		return errors.New("no entries found")
	}
	transaction, ok := userEntries[transactionId]
	if !ok {
		return ledger.ErrEntryNotFound
	}

	if transaction.Amount == *amount && amount != nil {
		reversedTxn, updatedTxn := transaction.UpdateAmount(*amount)
		s.Save(reversedTxn)
		s.Save(updatedTxn)
		return nil
	}

	return nil
}

func (s *LedgerMemoryStore) List(userId uuid.UUID) ([]*ledger.Entry, error) {
	var result []*ledger.Entry
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
