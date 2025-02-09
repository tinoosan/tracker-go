package ledger

import (
	"errors"
	"time"
	vo "trackergo/internal/domain/valueobjects"

	"github.com/google/uuid"
	)

type Account struct {
	Details      *vo.AccountDetails
	UserID       uuid.UUID
	TotalDebits  *vo.Money
	TotalCredits *vo.Money
	IsActive     bool
	CreatedAt    time.Time
}

func NewAccount(details *vo.AccountDetails, userID uuid.UUID, currency vo.Currency) *Account {
	return &Account{
		Details:   details,
    UserID: userID,
    TotalDebits: &vo.Money{Amount: 0, Currency: currency},
    TotalCredits: &vo.Money{Amount: 0, Currency: currency},
		IsActive:  true,
		CreatedAt: time.Now(),
	}
}

func (a *Account) GetTotalDebits() *vo.Money {
	return a.TotalDebits
}

func (a *Account) GetTotalCredits() *vo.Money {
	return a.TotalCredits
}

func (a *Account) Debit(money *vo.Money) error {
  if money.Currency != a.TotalDebits.Currency {
    return errors.New("currency mismatch for debit")
  }
  totalDebits, err := a.TotalDebits.Add(money)
  if err != nil {
    return err
  }

  a.TotalDebits = totalDebits
  return nil
}

func (a *Account) Credit(money *vo.Money) error {
  if money.Currency != a.TotalCredits.Currency {
    return errors.New("currency mismatch for credit")
  }
  totalCredits, err := a.TotalCredits.Add(money)
  if err != nil {
    return err
  }
  a.TotalCredits = totalCredits
  return nil
}

func (a *Account) CurrentBalance() (*vo.Money, error) {

  balance, err := a.TotalDebits.Subtract(a.TotalCredits)
  if err != nil {
    return &vo.Money{}, nil
  }
	return balance, nil
}
