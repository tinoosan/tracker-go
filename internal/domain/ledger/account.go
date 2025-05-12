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
	CreatedAt    *vo.DateTime
  UpdatedAt    *vo.DateTime
}


// NewAccount constructs a new Account instance with the provided details, user ID, and currency.
// It initializes debit and credit totals to zero, sets the account as active, and assigns current timestamps.
func NewAccount(details *vo.AccountDetails, userID uuid.UUID, currency vo.Currency) *Account {
	now := vo.NewDateTime(time.Now())
	return &Account{
		Details:   details,
		UserID: userID,
		TotalDebits: &vo.Money{Amount: 0, Currency: currency},
		TotalCredits: &vo.Money{Amount: 0, Currency: currency},
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (a *Account) GetTotalDebits() *vo.Money {
	return a.TotalDebits
}

func (a *Account) GetTotalCredits() *vo.Money {
	return a.TotalCredits
}

func (a *Account) Debit(money *vo.Money) error {
  now := vo.NewDateTime(time.Now())
  if money.Currency != a.TotalDebits.Currency {
    return errors.New("currency mismatch for debit")
  }
  totalDebits, err := a.TotalDebits.Add(money)
  if err != nil {
    return err
  }

  a.TotalDebits = totalDebits
  a.UpdatedAt = now
  return nil
}

func (a *Account) Credit(money *vo.Money) error {
  now := vo.NewDateTime(time.Now())
  if money.Currency != a.TotalCredits.Currency {
    return errors.New("currency mismatch for credit")
  }
  totalCredits, err := a.TotalCredits.Add(money)
  if err != nil {
    return err
  }
  a.TotalCredits = totalCredits
  a.UpdatedAt = now
  return nil
}

func (a *Account) CurrentBalance() (*vo.Money, error) {

  balance, err := a.TotalDebits.Subtract(a.TotalCredits)
  if err != nil {
    return &vo.Money{}, nil
  }
	return balance, nil
}
