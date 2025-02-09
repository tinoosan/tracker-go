package ledger

import (
	"fmt"
	"testing"
	"time"
	vo "trackergo/internal/domain/valueobjects"

	"github.com/google/uuid"
)

var (
	money = &vo.Money{
		Amount:   1000,
		Currency: vo.SupportedCurrencies["GBP"],
	}

	entry = &Entry{
		ID:             uuid.New(),
		PrimaryAccCode: vo.Code(101),
		LinkedAccCode:  vo.Code(601),
		UserID:         uuid.New(),
		EntryType:      Debit,
		Money:          money,
		Description:    "test",
		LinkedTxnID:    uuid.New(),
		CreatedAt:      vo.NewDateTime(time.Now()),
		UpdatedAt:      vo.NewDateTime(time.Now()),
	}
)

func Test_GetBalance(t *testing.T) {
	tests := []struct {
		name      string
		entry1    *Entry
		expected  float64
		expectErr bool
	}{
		{
			name:      "test GetBalance()",
			entry1:    entry,
			expected:  10.00,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			result := tt.entry1.GetBalance()

			if result != float64(tt.entry1.Money.Amount)/100 {
				fmt.Errorf("expected %.2f but got %.2f", float64(tt.entry1.Money.Amount)/100, result)
			}

		})
	}
}
