package application

import (
	"errors"
	"trackergo/internal/domain/ledger"
  vo "trackergo/internal/domain/valueobjects"

	"github.com/google/uuid"
)

type LedgerService struct {
	repo       LedgerRepository
	accService *AccountService
}

func NewLedgerService(repo LedgerRepository, accService *AccountService) *LedgerService {
	return &LedgerService{repo: repo, accService: accService}
}

func (s *LedgerService) CreateTransaction(debitName, creditName string, userID uuid.UUID, amount float64, currency string, description string) (*ledger.Entry, *ledger.Entry, error) {

	debitAccount, err := s.accService.GetAccountByName(debitName, userID)
	if err != nil {
		return nil, nil, err
	}
	creditAccount, err := s.accService.GetAccountByName(creditName, userID)
	if err != nil {
		return nil, nil, err
	}
	debitTxn, err := createDebitEntry(debitAccount.Details.Code, creditAccount.Details.Code, userID, amount, currency, description)
	if err != nil {
		return &ledger.Entry{}, &ledger.Entry{}, nil
	}
	creditTxn, err := createCreditEntry(creditAccount.Details.Code, debitAccount.Details.Code, userID, amount, currency, description)
	if err != nil {
		return &ledger.Entry{}, &ledger.Entry{}, nil
	}

	debitTxn.LinkedTxnID = creditTxn.ID
	creditTxn.LinkedTxnID = debitTxn.ID

	debitTxn.Process()
	creditTxn.Process()

	debitAccount.Debit(amount)
	creditAccount.Credit(amount)

	s.repo.Save(debitTxn)
	s.repo.Save(creditTxn)

	return debitTxn, creditTxn, nil
}

func (s *LedgerService) ReverseEntry(txID, userID uuid.UUID) error {
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
  amount float64, currency string, description string) (*ledger.Entry, error) {
	debitTxn, err := ledger.NewEntry(primaryAccID, linkedAccID, userID, ledger.Debit, amount, currency, description)
	if err != nil {
		return &ledger.Entry{}, err
	}
	return debitTxn, nil
}

func createCreditEntry(primaryAccID, linkedAccID vo.Code, userID uuid.UUID, 
  amount float64, currency string, description string) (*ledger.Entry, error) {
	creditTxn, err := ledger.NewEntry(primaryAccID, linkedAccID, userID, ledger.Credit, amount, currency, description)
	if err != nil {
		return &ledger.Entry{}, nil
	}
	return creditTxn, nil
}
