package accounts

import (
	"fmt"
	"github.com/google/uuid"
)



type Service struct {
	repo       AccountRepository
}

func NewService(repo AccountRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateAccount(userID uuid.UUID, name string, accountType Type) (*Account, error) {
  account := NewAccount(userID, name, accountType)

	err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *Service) CreateDefaultAccounts(userID uuid.UUID) error {

defaultAccounts := map[string]string{
		"Cash":           "ASSET",
		"Rent":           "EXPENSE",
		"Utilities":      "EXPENSE",
		"Groceries":      "EXPENSE",
		"Personal":       "EXPENSE",
		"Transportation": "EXPENSE",
		"Entertainment":  "EXPENSE",
		"Subscriptions":  "EXPENSE",
		"Miscellaneous":  "EXPENSE",
	}
	fmt.Println("Creating accounts...")
	for k, v := range defaultAccounts {
		newAccount, err := s.CreateAccount(userID, k, Type(v))
		if err != nil {
			return err
		}
		fmt.Printf("Account %v with ID %v has been created\n", newAccount.Name, newAccount.Id)
	}
	fmt.Println("Accounts successfully created")
  return nil
}

func (s *Service) GetAccountByID(accountID, userID uuid.UUID) (*Account, error) {
  account, err := s.repo.FindByID(accountID, userID)
  if err != nil {
    return nil, err
  }
  return account, nil
}

func (s *Service) GetAccountByName(name string, userID uuid.UUID) (*Account, error) {
  account , err := s.repo.FindByName(userID, name)
  if err != nil || account == nil {
    return nil, err
  }

  return account, nil
}

func (s *Service) UpdateAccount(accountID, userID uuid.UUID, name string) error {
  err := s.repo.Update(accountID, userID, name)
  if err != nil {
    return err
  }
  return nil
}

func (s *Service) DeleteAccount(accountID, userID uuid.UUID) error {
  err := s.repo.Delete(accountID, userID)
  if err != nil {
    return err
  }
  return nil
}

