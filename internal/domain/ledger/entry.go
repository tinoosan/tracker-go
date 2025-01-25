package ledger

import (
	"time"

	"github.com/google/uuid"
)

type EntryType string

const (
	Debit  EntryType = "Debit"
	Credit EntryType = "Credit"
)

type Type string

const (
	TypeUnknown        Type = "UNKNOWN"
	TypeSalaryIncome   Type = "SALARY_INCOME"
	TypeRentPayment    Type = "RENT_PAYMENT"
	TypeUtilities      Type = "UTILITIES"
	TypeGroceries      Type = "GROCERIES"
	TypePersonal       Type = "PERSONAL"
	TypeTransportation Type = "TRANSPORTATION"
	TypeEntertainment  Type = "ENTERTAINMENT"
	TypeSubscriptions  Type = "SUBSCRIPTIONS"
	TypeMiscellaneous  Type = "MISC"
)

type Entry struct {
	ID             uuid.UUID
	PrimaryAccCode Code
	LinkedAccCode  Code
	UserID         uuid.UUID
	EntryType      EntryType
	Money          *Money
	Description    string
	LinkedTxnID    uuid.UUID
	Reversal       bool
	ReversalOf     uuid.UUID
	Processed      bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewEntry(primaryAccCode, linkedAccCode Code, userID uuid.UUID,
	entryType EntryType, amount float64, currency Currency, description string) (*Entry, error) {
  money, err := NewMoney(amount, currency)
  if err != nil {
    return &Entry{}, err
  }
	return &Entry{
		ID:             uuid.New(),
		PrimaryAccCode: primaryAccCode,
		LinkedAccCode:  linkedAccCode,
		UserID:         userID,
		EntryType:      entryType,
		Money:         money,
		Description:    description,
		CreatedAt:      time.Now(),
	}, nil

}

func (e *Entry) GetBalance() float64 {
	if e.EntryType == Debit {
		return float64(e.Money.Amount)
	}
	return -float64(e.Money.Amount)
}

func (t *Entry) Process() {
	t.Processed = true
	t.UpdatedAt = time.Now()
}

func (t *Entry) Reverse() (*Entry, error) {
	reversedTxn, err := NewEntry(t.PrimaryAccCode,
		t.LinkedAccCode,
		t.UserID,
		t.EntryType.reverseOf(),
		float64(t.Money.Amount),
    t.Money.Currency,
		t.Description)

  if err != nil {
    return &Entry{}, err
  }

	reversedTxn.ReversalOf = t.ID
	reversedTxn.Reversal = true

	return reversedTxn, nil
}

func (t *Entry) UpdateAmount(amount float64, currency Currency) (*Entry, *Entry, error) {
	reversedTxn, err := t.Reverse()

  if err != nil {
    return &Entry{}, &Entry{}, err
  }

	updatedTxn, err := NewEntry(reversedTxn.PrimaryAccCode, reversedTxn.LinkedAccCode,
		reversedTxn.UserID, t.EntryType, amount, currency, t.Description)

  if err != nil {
    return &Entry{}, &Entry{}, err
  }

	return reversedTxn, updatedTxn, nil
}

func (e EntryType) reverseOf() EntryType {
	rules := map[EntryType]EntryType{
		Debit:  Credit,
		Credit: Debit,
	}
	for k, v := range rules {
		if k == e {
			return v
		}
	}
	return ""
}
