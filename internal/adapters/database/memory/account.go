package memory

import (
	"errors"
	"trackergo/internal/domain/ledger"

	"github.com/google/uuid"
)

type AccountMemoryStore struct {
	// Takes UserID and AccountID
	UserAccounts map[uuid.UUID]map[ledger.Code]*ledger.Account
}

func NewAccountMemoryStore() *AccountMemoryStore {
	return &AccountMemoryStore{UserAccounts: make(map[uuid.UUID]map[ledger.Code]*ledger.Account)}
}

func (s *AccountMemoryStore) Save(account *ledger.Account) error {
	userAccounts, exists := s.UserAccounts[account.UserID]
	if !exists {
		s.UserAccounts[account.UserID] = make(map[ledger.Code]*ledger.Account)
		userAccounts = s.UserAccounts[account.UserID]
	}
	if account == nil {
		return errors.New("Account is nil")
	}
	userAccounts[account.Code] = account
	return nil
}

func (s *AccountMemoryStore) FindByCode(code ledger.Code, userID uuid.UUID) (*ledger.Account, error) {
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

func (s *AccountMemoryStore) FindByName(userID uuid.UUID, name string) (*ledger.Account, error) {
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

func (s *AccountMemoryStore) List(userID uuid.UUID) ([]*ledger.Account, error) {
	var result []*ledger.Account
	userAccounts, ok := s.UserAccounts[userID]
	if !ok {
		return result, errors.New("No accounts exists for user")
	}
	for _, account := range userAccounts {
		result = append(result, account)
	}
	return result, nil
}

func (s *AccountMemoryStore) Update(code ledger.Code, userID uuid.UUID, name string) error {
	account, err := s.FindByCode(code, userID)
	if err != nil {
		return err
	}
	if account.Name != "" || account.Name != name {
		account.Name = name
	}
	return nil
}

func (s *AccountMemoryStore) Delete(code ledger.Code, userID uuid.UUID) error {

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
