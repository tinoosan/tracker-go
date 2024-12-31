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
	ID               uuid.UUID
	PrimaryAccountID uuid.UUID
	LinkedAccountID  uuid.UUID
	UserID           uuid.UUID
	EntryType        EntryType
	Amount           float64
	Description      string
	LinkedTxnID      uuid.UUID
	Reversal         bool
	ReversalOf       uuid.UUID
	Processed        bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewEntry(primaryAccID, linkedAccID, userID uuid.UUID,
	entryType EntryType, amount float64, description string) *Entry {
	return &Entry{
		ID:               uuid.New(),
		PrimaryAccountID: primaryAccID,
		LinkedAccountID:  linkedAccID,
		UserID:           userID,
		EntryType:        entryType,
		Amount:           amount,
		CreatedAt:        time.Now(),
	}

}

func (e *Entry) GetBalance() float64 {
	if e.EntryType == Debit {
		return e.Amount
	}
	return -e.Amount
}

func (t *Entry) Process() {
	t.Processed = true
	t.UpdatedAt = time.Now()
}

func (t *Entry) Reverse() *Entry {
	reversedTxn := NewEntry(t.PrimaryAccountID,
		t.LinkedAccountID,
		t.UserID,
		t.EntryType.reverseOf(),
		t.Amount,
		t.Description)

	reversedTxn.ReversalOf = t.ID
	reversedTxn.Reversal = true

	return reversedTxn
}

func (t *Entry) UpdateAmount(amount float64) (*Entry, *Entry) {
	reversedTxn := t.Reverse()
	updatedTxn := NewEntry(reversedTxn.PrimaryAccountID,
		reversedTxn.LinkedAccountID,
		reversedTxn.UserID, t.EntryType, amount, t.Description)

	return reversedTxn, updatedTxn
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
