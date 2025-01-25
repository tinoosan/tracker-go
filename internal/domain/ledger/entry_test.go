package ledger_test

import (
	"fmt"
	"testing"
	"time"
	"trackergo/internal/domain/ledger"

	"github.com/google/uuid"
)

var (
	money = &ledger.Money{
		Amount:   1000,
		Currency: ledger.GBP,
	}

	entry = &ledger.Entry{
		ID:             uuid.New(),
		PrimaryAccCode: ledger.Code(101),
		LinkedAccCode:  ledger.Code(601),
		UserID:         uuid.New(),
		EntryType:      ledger.Debit,
		Money:          money,
		Description:    "test",
		LinkedTxnID:    uuid.New(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
)

func Test_GetBalance(t *testing.T) {
	tests := []struct {
		name      string
		entry1    *ledger.Entry
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
