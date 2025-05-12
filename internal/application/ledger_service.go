package application

import (
	"errors"
	"sync"
	"trackergo/internal/domain/ledger"
	vo "trackergo/internal/domain/valueobjects"

	"github.com/google/uuid"
)

type LedgerService struct {
	repo       LedgerRepository
	accService *AccountService
  exchangeRateProvider ExchangeRateProvider
  mutex sync.Mutex
}

func NewLedgerService(repo LedgerRepository, accService *AccountService, provider ExchangeRateProvider) *LedgerService {
	return &LedgerService{repo: repo, accService: accService, exchangeRateProvider: provider}
}

func (s *LedgerService) CreateTransaction(debitName, creditName string, userID uuid.UUID, 
  amount float64, currency string, description string) (*ledger.Entry, *ledger.Entry, error) {
  
  s.mutex.Lock()
  defer s.mutex.Unlock()

  money, err := vo.NewMoney(amount, currency)
  if err != nil {
    return &ledger.Entry{}, &ledger.Entry{}, err
  }

	debitAccount, err := s.accService.GetAccountByName(debitName, userID)
	if err != nil {
		return nil, nil, err
	}

  exchangeRate, err := vo.NewRatio(1.0)
  if err != nil {
    return nil, nil, err
  }

  if money.Currency != debitAccount.TotalDebits.Currency {
    exchangeRate, err = s.exchangeRateProvider.GetExchangeRate(money.Currency.Code, debitAccount.TotalDebits.Currency.Code)
    if err != nil {
      return nil, nil, err
    }
  }

  if money.Currency.Code != debitAccount.TotalDebits.Currency.Code {
    money, err = money.Convert(debitAccount.TotalDebits.Currency.Code, exchangeRate)
    if err != nil {
      return nil, nil, err
    }
  }

	creditAccount, err := s.accService.GetAccountByName(creditName, userID)
	if err != nil {
		return nil, nil, err
	}
	debitTxn, err := createDebitEntry(debitAccount.Details.Code, creditAccount.Details.Code, userID, money, description)
	if err != nil {
		return &ledger.Entry{}, &ledger.Entry{}, nil
	}
	creditTxn, err := createCreditEntry(creditAccount.Details.Code, debitAccount.Details.Code, userID, money, description)
	if err != nil {
		return &ledger.Entry{}, &ledger.Entry{}, nil
	}

	debitTxn.LinkedTxnID = creditTxn.ID
	creditTxn.LinkedTxnID = debitTxn.ID

	debitTxn.Process()
	creditTxn.Process()

	err = debitAccount.Debit(money)
  if err != nil {
    return &ledger.Entry{}, &ledger.Entry{}, err
  }

	err = creditAccount.Credit(money)
  if err != nil {
    return &ledger.Entry{}, &ledger.Entry{}, err
  }

	s.repo.Save(debitTxn)
	s.repo.Save(creditTxn)

	return debitTxn, creditTxn, nil
}

func (s *LedgerService) ReverseEntry(txID, userID uuid.UUID) error {
  s.mutex.Lock()
  defer s.mutex.Unlock()

	tx, err := s.repo.FindByID(txID, userID)
	if err != nil {
		return err
	}
	tx.Reverse()

	linkedTx, err := s.repo.FindByID(tx.LinkedTxnID, userID)
	if err != nil {
		return err
	}

	linkedTx.Reverse()
	linkedTx.Process()
	return nil
}

func (s *LedgerService) GetTAccount(name string, userID uuid.UUID) ([]*ledger.Entry, *ledger.Account, error) {
  s.mutex.Lock()
  defer s.mutex.Unlock()
	var result []*ledger.Entry

	account, err := s.accService.GetAccountByName(name, userID)
	if err != nil {
		return nil, nil, err
	}

	userEntries, err := s.repo.List(userID)
	if err != nil {
		return nil, account, err
	}

	if userEntries == nil {
		return nil, nil, errors.New("no journal entries found")
	}

	for _, v := range userEntries {
		if v.PrimaryAccCode == account.Details.Code {
			result = append(result, v)
		}
	}

	return result, account, nil
}

func createDebitEntry(primaryAccID, linkedAccID vo.Code, userID uuid.UUID, 
  money *vo.Money, description string) (*ledger.Entry, error) {
	debitTxn, err := ledger.NewEntry(primaryAccID, linkedAccID, userID, ledger.Debit, money, description)
	if err != nil {
		return &ledger.Entry{}, err
	}
	return debitTxn, nil
}

func createCreditEntry(primaryAccID, linkedAccID vo.Code, userID uuid.UUID, 
  money *vo.Money, description string) (*ledger.Entry, error) {
	creditTxn, err := ledger.NewEntry(primaryAccID, linkedAccID, userID, ledger.Credit, money, description)
	if err != nil {
		return &ledger.Entry{}, nil
	}
	return creditTxn, nil
}
