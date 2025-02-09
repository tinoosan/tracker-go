package application

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"trackergo/internal/domain/ledger"
	vo "trackergo/internal/domain/valueobjects"

	"github.com/google/uuid"
)

type AccountService struct {
	repo      AccountRepository
	codeIndex map[vo.AccountType]int
	mutex     sync.Mutex
}

func NewAccountService(repo AccountRepository) *AccountService {
	return &AccountService{
		repo:      repo,
		codeIndex: make(map[vo.AccountType]int),
	}
}

func (s *AccountService) CreateAccount(userID uuid.UUID, name string, accountType vo.AccountType, currency vo.Currency) (*ledger.Account, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if userID.String() == "00000000-0000-0000-0000-000000000000" {
		return nil, errors.New("invalid userID format")
	}

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	name = strings.ToUpper(name)

	if _, ok := map[vo.AccountType]bool{
		vo.TypeAsset:     true,
		vo.TypeLiability: true,
		vo.TypeEquity:    true,
		vo.TypeExpense:   true,
		vo.TypeRevenue:   true,
	}[accountType]; !ok {
		return nil, errors.New("invalid account type")
	}

	if _, exists := s.codeIndex[accountType]; !exists {
		s.codeIndex[accountType] = int(getBaseCode(accountType))
	}

	code := s.codeIndex[accountType]
	s.codeIndex[accountType]++

	details, err := vo.NewAccountDetails(vo.Code(code), name, accountType)
	if err != nil {
		return &ledger.Account{}, err
	}
	account := ledger.NewAccount(details, userID, currency)

	err = s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountService) CreateDefaultAccounts(userID uuid.UUID, currency string) error {

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
		newAccount, err := s.CreateAccount(userID, k, vo.AccountType(v), vo.SupportedCurrencies[currency])
		if err != nil {
			return err
		}
		fmt.Printf("Account %v with code %v has been created\n", newAccount.Details.Name, newAccount.Details.Code)
	}
	fmt.Println("Accounts successfully created")
	return nil
}

func (s *AccountService) GetAccountByID(code vo.Code, userID uuid.UUID) (*ledger.Account, error) {
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

func (s *AccountService) UpdateAccount(code vo.Code, userID uuid.UUID, name string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	err := s.repo.Update(code, userID, name)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountService) DeleteAccount(code vo.Code, userID uuid.UUID) error {
	err := s.repo.Delete(code, userID)
	if err != nil {
		return err
	}
	return nil
}

func getBaseCode(accountType vo.AccountType) vo.Code {
	switch accountType {
	case vo.TypeAsset:
		return vo.CodeAsset
	case vo.TypeLiability:
		return vo.CodeLiability
	case vo.TypeEquity:
		return vo.CodeEquity
	case vo.TypeExpense:
		return vo.CodeExpense
	case vo.TypeRevenue:
		return vo.CodeRevenue
	default:
		return 0

	}
}
