package ledger

import (
	"trackergo/internal/accounts"

	"github.com/google/uuid"
)

type Service struct {
	repo       LedgerRepository
	accService accounts.Service
}

func NewService(repo LedgerRepository, accService accounts.Service) *Service {
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
	debitTxn := createDebitEntry(debitAccount.Id, creditAccount.Id, userID, amount, description)
	creditTxn := createCreditEntry(creditAccount.Id, debitAccount.Id, userID, amount, description)

	debitTxn.LinkedTxnID = creditTxn.PrimaryAccountID
	creditTxn.LinkedTxnID = debitTxn.PrimaryAccountID

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
	userEntries, err := s.repo.List(userID)
	if err != nil {
		return userEntries, nil, err
	}

	account, err := s.accService.GetAccountByName(name, userID)
  if err != nil {
    return userEntries, nil, err
  }

	for _, v := range userEntries {
    if v.PrimaryAccountID == account.Id {
      result = append(result, v)
    }
	}

  return result, account, nil
}

func createDebitEntry(primaryAccID, linkedAccID, userID uuid.UUID, amount float64, description string) *Entry {
	debitTxn := NewEntry(primaryAccID, linkedAccID, userID, Debit, amount, description)
	return debitTxn
}

func createCreditEntry(primaryAccID, linkedAccID, userID uuid.UUID, amount float64, description string) *Entry {
	creditTxn := NewEntry(primaryAccID, linkedAccID, userID, Credit, amount, description)
	return creditTxn
}
