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
	totalDebits  float64
	totalCredits float64
	IsActive     bool
	CreatedAt    time.Time
}

func NewAccount(details *vo.AccountDetails, userID uuid.UUID) *Account {
	return &Account{
		Details:   details,
    UserID: userID,
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
