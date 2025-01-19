package application

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"trackergo/internal/domain/ledger"

	"github.com/google/uuid"
)

type AccountService struct {
	repo      AccountRepository
	codeIndex map[ledger.AccountType]int
	mutex     sync.Mutex
}

func NewAccountService(repo AccountRepository) *AccountService {
	return &AccountService{
		repo:      repo,
		codeIndex: make(map[ledger.AccountType]int),
	}
}

func (s *AccountService) CreateAccount(userID uuid.UUID, name string, accountType ledger.AccountType) (*ledger.Account, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if userID.String() == "00000000-0000-0000-0000-000000000000" {
		return nil, errors.New("invalid userID format")
	}

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	name = strings.ToUpper(name)

	if _, ok := map[ledger.AccountType]bool{
		ledger.TypeAsset:     true,
		ledger.TypeLiability: true,
		ledger.TypeEquity:    true,
		ledger.TypeExpense:   true,
		ledger.TypeRevenue:   true,
	}[accountType]; !ok {
		return nil, errors.New("invalid account type")
	}

	if _, exists := s.codeIndex[accountType]; !exists {
		s.codeIndex[accountType] = int(getBaseCode(accountType))
	}

	code := s.codeIndex[accountType]
	s.codeIndex[accountType]++
	account := ledger.NewAccount(ledger.Code(code), userID, name, accountType)

	err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountService) CreateDefaultAccounts(userID uuid.UUID) error {

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
		newAccount, err := s.CreateAccount(userID, k, ledger.AccountType(v))
		if err != nil {
			return err
		}
		fmt.Printf("Account %v with code %v has been created\n", newAccount.Name, newAccount.Code)
	}
	fmt.Println("Accounts successfully created")
	return nil
}

func (s *AccountService) GetAccountByID(code ledger.Code, userID uuid.UUID) (*ledger.Account, error) {
	account, err := s.repo.FindByCode(code, userID)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AccountService) GetAccountByName(name string, userID uuid.UUID) (*ledger.Account, error) {
	account, err := s.repo.FindByName(userID, name)
	if err != nil || account == nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountService) GetChartOfAccounts(userID uuid.UUID) ([]*ledger.Account, error) {
	accounts, err := s.repo.List(userID)
	if err != nil {
		return accounts, nil
	}
	return accounts, nil
}

func (s *AccountService) UpdateAccount(code ledger.Code, userID uuid.UUID, name string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	err := s.repo.Update(code, userID, name)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountService) DeleteAccount(code ledger.Code, userID uuid.UUID) error {
	err := s.repo.Delete(code, userID)
	if err != nil {
		return err
	}
	return nil
}

func getBaseCode(accountType ledger.AccountType) ledger.Code {
	switch accountType {
	case ledger.TypeAsset:
		return ledger.CodeAsset
	case ledger.TypeLiability:
		return ledger.CodeLiability
	case ledger.TypeEquity:
		return ledger.CodeEquity
	case ledger.TypeExpense:
		return ledger.CodeExpense
	case ledger.TypeRevenue:
		return ledger.CodeRevenue
	default:
		return 0

	}
}
