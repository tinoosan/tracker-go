package accounts

import (
	"errors"

	"github.com/google/uuid"
)

var (
	_ AccountRepository = &InMemoryStore{}
)

type AccountGetter interface {
	FindByID(accountID, userID uuid.UUID) (*Account, error)
	FindByName(userID uuid.UUID, name string) (*Account, error)
}

type AccountRepository interface {
  AccountGetter
	Save(account *Account) error
	Update(accountID, userID uuid.UUID, name string) error
	Delete(accountID, userID uuid.UUID) error
}

type InMemoryStore struct {
	// Takes UserID and AccountID
	UserAccounts map[uuid.UUID]map[uuid.UUID]*Account
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{UserAccounts: make(map[uuid.UUID]map[uuid.UUID]*Account)}
}

func (s *InMemoryStore) Save(account *Account) error {
	userAccounts, exists := s.UserAccounts[account.UserID]
	if !exists {
		s.UserAccounts[account.UserID] = make(map[uuid.UUID]*Account)
		userAccounts = s.UserAccounts[account.UserID]
	}
	if account == nil {
		return errors.New("Account is nil")
	}
	userAccounts[account.Id] = account
	return nil
}

func (s *InMemoryStore) FindByID(accountID, userID uuid.UUID) (*Account, error) {
	userAccounts, ok := s.UserAccounts[userID]
	if !ok {
		return nil, errors.New("No accounts found")
	}

	account, ok := userAccounts[accountID]
	if !ok {
		return nil, errors.New("Account could not be found")
	}
	return account, nil
}

func (s *InMemoryStore) FindByName(userID uuid.UUID, name string) (*Account, error) {
	userAccounts, ok := s.UserAccounts[userID]
	if !ok {
		return nil, errors.New("No accounts found")
	}

	for _, account := range userAccounts {
		if account.Name == name {
			return account, nil
		}
	}
	return nil, errors.New("No account with name %v. Please try again.\n")
}

func (s *InMemoryStore) Update(accountID, userID uuid.UUID, name string) error {
	account, err := s.FindByID(accountID, userID)
	if err != nil {
		return err
	}
	if account.Name != "" || account.Name != name {
		account.Name = name
	}
	return nil
}

func (s *InMemoryStore) Delete(accountID, userID uuid.UUID) error {

	account, err := s.FindByID(accountID, userID)
	if err != nil {
		return err
	}
	if account.IsActive == false {
		return errors.New("Account is already deactivated")
	}
	account.IsActive = false
	return nil
}
