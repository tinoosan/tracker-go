package ledger

import (
	"errors"
	"trackergo/internal/accounts"

	"github.com/google/uuid"
)

type Service struct {
	repo       LedgerRepository
	accService *accounts.Service
}

func NewService(repo LedgerRepository, accService *accounts.Service) *Service {
	return &Service{repo: repo, accService: accService}
}

func (s *Service) CreateTransaction(debitName, creditName string, userID uuid.UUID, amount float64, description string) (*Entry, *Entry, error) {

	debitAccount, err := s.accService.GetAccountByName(debitName, userID)
	if err != nil {
		return nil, nil, err
	}
	creditAccount, err := s.accService.GetAccountByName(creditName, userID)
	if err != nil {
		return nil, nil, err
	}
	debitTxn := createDebitEntry(debitAccount.Code, creditAccount.Code, userID, amount, description)
	creditTxn := createCreditEntry(creditAccount.Code, debitAccount.Code, userID, amount, description)

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

func (s *Service) ReverseEntry(txID, userID uuid.UUID) error {
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

func (s *Service) GetTAccount(name string, userID uuid.UUID) ([]*Entry, *accounts.Account, error) {
	var result []*Entry

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
		if v.PrimaryAccCode == account.Code {
			result = append(result, v)
		}
	}

	return result, account, nil
}

func createDebitEntry(primaryAccID, linkedAccID accounts.Code, userID uuid.UUID, amount float64, description string) *Entry {
	debitTxn := NewEntry(primaryAccID, linkedAccID, userID, Debit, amount, description)
	return debitTxn
}

func createCreditEntry(primaryAccID, linkedAccID accounts.Code, userID uuid.UUID, amount float64, description string) *Entry {
	creditTxn := NewEntry(primaryAccID, linkedAccID, userID, Credit, amount, description)
	return creditTxn
}
