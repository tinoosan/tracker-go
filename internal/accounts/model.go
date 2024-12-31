package accounts

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id           uuid.UUID
	UserID       uuid.UUID
	Name         string
	Type         Type
	totalDebits  float64
	totalCredits float64
	IsActive     bool
	CreatedAt    time.Time
}

type Type string

const (
	TypeAsset     Type = "ASSET"
	TypeLiability Type = "LIABILITY"
	TypeEquity    Type = "EQUITY"
	TypeExpense   Type = "EXPENSE"
	TypeRevenue   Type = "REVENUE"
)

func NewAccount(userID uuid.UUID, name string, accountType Type) *Account {
	return &Account{
		Id:        uuid.New(),
		UserID:    userID,
		Name:      name,
		Type:      accountType,
		IsActive:  true,
		CreatedAt: time.Now(),
	}
}

func (a *Account) GetTotalDebits() float64 {
  return a.totalDebits
}

func (a *Account) GetTotalCredits() float64 {
  return a.totalCredits
}

func (a *Account) Debit(amount float64) error {
	if amount < 0 {
    return errors.New("amount must be positive")
  }
  a.totalDebits += amount
	return nil
}

func (a *Account) Credit(amount float64) error {
	if amount < 0 {
    return errors.New("amount must be positive")
  }
  a.totalCredits += amount

	return nil
}

func (a *Account) CurrentBalance() float64 {
  return a.totalDebits - a.totalCredits
}
