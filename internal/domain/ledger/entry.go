package ledger

import (
	"time"
	vo "trackergo/internal/domain/valueobjects"

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
	PrimaryAccCode vo.Code
	LinkedAccCode  vo.Code
	UserID         uuid.UUID
	EntryType      EntryType
	Money          *vo.Money
	Description    string
	LinkedTxnID    uuid.UUID
	Reversal       bool
	ReversalOf     uuid.UUID
	Processed      bool
	CreatedAt      *vo.DateTime
	UpdatedAt      *vo.DateTime
}

func NewEntry(primaryAccCode, linkedAccCode vo.Code, userID uuid.UUID,
	entryType EntryType, money *vo.Money, description string) (*Entry, error) {
	now := vo.NewDateTime(time.Now())
	return &Entry{
		ID:             uuid.New(),
		PrimaryAccCode: primaryAccCode,
		LinkedAccCode:  linkedAccCode,
		UserID:         userID,
		EntryType:      entryType,
		Money:          money,
		Description:    description,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil

}

func (e *Entry) GetBalance() float64 {
	if e.EntryType == Debit {
		return e.Money.GetAmount()
	}
	return -e.Money.GetAmount()
}

func (t *Entry) Process() {
  now := vo.NewDateTime(time.Now())
	t.Processed = true
	t.UpdatedAt = now

}

func (t *Entry) Reverse() (*Entry, error) {
	reversedTxn, err := NewEntry(t.PrimaryAccCode,
		t.LinkedAccCode,
		t.UserID,
		t.EntryType.reverseOf(),
		t.Money,
		t.Description)

	if err != nil {
		return &Entry{}, err
	}

	reversedTxn.ReversalOf = t.ID
	reversedTxn.Reversal = true

	return reversedTxn, nil
}

func (t *Entry) UpdateAmount(amount float64) (*Entry, *Entry, error) {
	reversedTxn, err := t.Reverse()

	if err != nil {
		return &Entry{}, &Entry{}, err
	}

	updatedTxn, err := NewEntry(reversedTxn.PrimaryAccCode, reversedTxn.LinkedAccCode,
		reversedTxn.UserID, t.EntryType, t.Money, t.Description)

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
