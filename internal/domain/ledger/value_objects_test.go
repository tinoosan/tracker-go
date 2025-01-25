package ledger_test

import (
	"testing"
	"trackergo/internal/domain/ledger"
)

func TestMoney_Add(t *testing.T) {
	tests := []struct {
		name      string
		money1    *ledger.Money
		money2    *ledger.Money
		expected  *ledger.Money
		expectErr bool
	}{
		{
			name: "Add same currency",
			money1: &ledger.Money{
				Amount:   1000, // £10.00
				Currency: ledger.GBP,
			},
			money2: &ledger.Money{
				Amount:   500, // £5.00
				Currency: ledger.GBP,
			},
			expected: &ledger.Money{
				Amount:   1500, // £15.00
				Currency: ledger.GBP,
			},
			expectErr: false,
		},
		{
			name: "Add different currencies",
			money1: &ledger.Money{
				Amount:   1000, // £10.00
				Currency: ledger.GBP,
			},
			money2: &ledger.Money{
				Amount:   500, // £5.00
				Currency: ledger.USD,
			},
			expected: &ledger.Money{},
			expectErr: true,
		},
	}

  for _, tt := range tests {
    t.Run(tt.name, func(*testing.T){
      result, err := tt.money1.Add(tt.money2)

      if tt.expectErr {
        if err == nil {
          t.Error("expected error, got nil")
        }
        return
      }

      if err != nil {
        t.Errorf("unexpected error: %v", err)
        return
      }

      if result.Amount != tt.expected.Amount || result.Currency != tt.expected.Currency {
        t.Errorf("expected %+v, but got %+v", tt.expected, result)
      }
    })
  }

}

func TestMoney_Subtract(t *testing.T) {
	tests := []struct {
		name      string
		money1    *ledger.Money
		money2    *ledger.Money
		expected  *ledger.Money
		expectErr bool
	}{
		{
			name: "Subtract same currency",
			money1: &ledger.Money{
				Amount:   1000, // £10.00
				Currency: ledger.GBP,
			},
			money2: &ledger.Money{
				Amount:   500, // £5.00
				Currency: ledger.GBP,
			},
			expected: &ledger.Money{
				Amount:   500, // £15.00
				Currency: ledger.GBP,
			},
			expectErr: false,
		},
		{
			name: "Subtract different currencies",
			money1: &ledger.Money{
				Amount:   1000, // £10.00
				Currency: ledger.GBP,
			},
			money2: &ledger.Money{
				Amount:   500, // £5.00
				Currency: ledger.USD,
			},
			expected: &ledger.Money{},
			expectErr: true,
		},
	}

  for _, tt := range tests {
    t.Run(tt.name, func(*testing.T){
      result, err := tt.money1.Subtract(tt.money2)

      if tt.expectErr {
        if err == nil {
          t.Error("expected error, got nil")
        }
        return
      }

      if err != nil {
        t.Errorf("unexpected error: %v", err)
        return
      }

      if result.Amount != tt.expected.Amount || result.Currency != tt.expected.Currency {
        t.Errorf("expected %+v, but got %+v", tt.expected, result)
      }
    })
  }

}


func TestMoney_Convert(t *testing.T) {
	tests := []struct {
		name      string
		money1    *ledger.Money
    exchangeRate float64
    targetCurrency ledger.Currency
		expected  *ledger.Money
		expectErr bool
	}{
		{
			name: "Convert to different currency",
			money1: &ledger.Money{
				Amount:   1000, // £10.00
				Currency: ledger.GBP,
			},
      exchangeRate: 0.60,
      targetCurrency: ledger.USD,
			expected: &ledger.Money{
				Amount:   600, // £15.00
				Currency: ledger.USD,
			},
			expectErr: false,
		},
		{
			name: "Convert to same currencies",
			money1: &ledger.Money{
				Amount:   1000, // £10.00
				Currency: ledger.GBP,
			},
      exchangeRate: 1.00,
      targetCurrency: ledger.GBP,
			expected: &ledger.Money{},
			expectErr: true,
		},
	}

  for _, tt := range tests {
    t.Run(tt.name, func(*testing.T){
      result, err := tt.money1.Convert(tt.targetCurrency, tt.exchangeRate)

      if tt.expectErr {
        if err == nil {
          t.Error("expected error, got nil")
        }
        return
      }

      if err != nil {
        t.Errorf("unexpected error: %v", err)
        return
      }

      if result.Amount != tt.expected.Amount || result.Currency != tt.expected.Currency {
        t.Errorf("expected %+v, but got %+v", tt.expected, result)
      }
    })
  }

}
