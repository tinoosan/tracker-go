package accounts

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type Service struct {
	repo      AccountRepository
	codeIndex map[Type]int
	mutex     sync.Mutex
}

func NewService(repo AccountRepository) *Service {
	return &Service{
		repo:      repo,
		codeIndex: make(map[Type]int),
	}
}

func (s *Service) CreateAccount(userID uuid.UUID, name string, accountType Type) (*Account, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

  if userID.String() == "00000000-0000-0000-0000-000000000000" {
    return nil, errors.New("invalid userID format")
  }

  if name == "" {
    return nil, errors.New("name cannot be empty")
  }

  name = strings.ToUpper(name)

	if _, ok := map[Type]bool{
		TypeAsset:     true,
		TypeLiability: true,
		TypeEquity:    true,
		TypeExpense:   true,
		TypeRevenue:   true,
	}[accountType]; !ok {
		return nil, errors.New("invalid account type")
	}

	if _, exists := s.codeIndex[accountType]; !exists {
		s.codeIndex[accountType] = int(getBaseCode(accountType))
	}

  code := s.codeIndex[accountType]
  s.codeIndex[accountType]++
	account := NewAccount(Code(code), userID, name, accountType)

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
		fmt.Printf("Account %v with code %v has been created\n", newAccount.Name, newAccount.Code)
	}
	fmt.Println("Accounts successfully created")
	return nil
}

func (s *Service) GetAccountByID(code Code, userID uuid.UUID) (*Account, error) {
  account, err := s.repo.FindByCode(code, userID)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *Service) GetAccountByName(name string, userID uuid.UUID) (*Account, error) {
	account, err := s.repo.FindByName(userID, name)
	if err != nil || account == nil {
		return nil, err
	}

	return account, nil
}

func (s *Service) GetChartOfAccounts(userID uuid.UUID) ([]*Account, error) {
  accounts, err := s.repo.List(userID)
  if err != nil {
    return accounts, nil
  }
  return accounts, nil
}

func (s *Service) UpdateAccount(code Code, userID uuid.UUID, name string) error {
  s.mutex.Lock()
  defer s.mutex.Unlock()
	err := s.repo.Update(code, userID, name)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteAccount(code Code, userID uuid.UUID) error {
	err := s.repo.Delete(code, userID)
	if err != nil {
		return err
	}
	return nil
}

func getBaseCode(accountType Type) Code {
	switch accountType {
	case TypeAsset:
		return CodeAsset
	case TypeLiability:
		return CodeLiability
	case TypeEquity:
		return CodeEquity
  case TypeExpense:
  return CodeExpense
    case TypeRevenue:
  return CodeRevenue
  default:
  return 0

	}
}
