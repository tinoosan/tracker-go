package accounts

import (
	"errors"

	"github.com/google/uuid"
)

var (
	_ AccountRepository = &InMemoryStore{}
)

type AccountGetter interface {
	FindByCode(code Code, userID uuid.UUID) (*Account, error)
	FindByName(userID uuid.UUID, name string) (*Account, error)
	List(userID uuid.UUID) ([]*Account, error)
}

type AccountRepository interface {
	AccountGetter
	Save(account *Account) error
	Update(code Code, userID uuid.UUID, name string) error
	Delete(code Code, userID uuid.UUID) error
}

type InMemoryStore struct {
	// Takes UserID and AccountID
	UserAccounts map[uuid.UUID]map[Code]*Account
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{UserAccounts: make(map[uuid.UUID]map[Code]*Account)}
}

func (s *InMemoryStore) Save(account *Account) error {
	userAccounts, exists := s.UserAccounts[account.UserID]
	if !exists {
		s.UserAccounts[account.UserID] = make(map[Code]*Account)
		userAccounts = s.UserAccounts[account.UserID]
	}
	if account == nil {
		return errors.New("Account is nil")
	}
	userAccounts[account.Code] = account
	return nil
}

func (s *InMemoryStore) FindByCode(code Code, userID uuid.UUID) (*Account, error) {
	userAccounts, ok := s.UserAccounts[userID]
	if !ok {
		return nil, errors.New("No accounts found")
	}

	account, ok := userAccounts[code]
	if !ok {
		return nil, errors.New("Account could not be found")
	}
	return account, nil
}

func (s *InMemoryStore) FindByName(userID uuid.UUID, name string) (*Account, error) {
	userAccounts, ok := s.UserAccounts[userID]
	if !ok {
		return nil, errors.New("No accounts exist for user")
	}

	for _, account := range userAccounts {
		if account.Name == name {
			return account, nil
		}
	}
	return nil, errors.New("No account with name %s. Please try again.\n")
}

func (s *InMemoryStore) List(userID uuid.UUID) ([]*Account, error) {
	var result []*Account
	userAccounts, ok := s.UserAccounts[userID]
	if !ok {
		return result, errors.New("No accounts exists for user")
	}
	for _, account := range userAccounts {
		result = append(result, account)
	}
	return result, nil
}

func (s *InMemoryStore) Update(code Code, userID uuid.UUID, name string) error {
	account, err := s.FindByCode(code, userID)
	if err != nil {
		return err
	}
	if account.Name != "" || account.Name != name {
		account.Name = name
	}
	return nil
}

func (s *InMemoryStore) Delete(code Code, userID uuid.UUID) error {

	account, err := s.FindByCode(code, userID)
	if err != nil {
		return err
	}
	if account.IsActive == false {
		return errors.New("Account is already deactivated")
	}
	account.IsActive = false
	return nil
}
