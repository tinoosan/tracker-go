package valueobjects

import (
	"testing"
)

var (
  exchangeRate, _ = NewRatio(0.60)
  exchangeRate2, _ = NewRatio(1.00)
)

func TestMoney_Add(t *testing.T) {
	tests := []struct {
		name      string
		money1    *Money
		money2    *Money
		expected  *Money
		expectErr bool
	}{
		{
			name: "Add same currency",
			money1: &Money{
				Amount:   1000, // £10.00
				Currency: SupportedCurrencies["GBP"],
			},
			money2: &Money{
				Amount:   500, // £5.00
				Currency: SupportedCurrencies["GBP"],
			},
			expected: &Money{
				Amount:   1500, // £15.00
				Currency: SupportedCurrencies["GBP"],
			},
			expectErr: false,
		},
		{
			name: "Add different currencies",
			money1: &Money{
				Amount:   1000, // £10.00
				Currency: SupportedCurrencies["GBP"],
			},
			money2: &Money{
				Amount:   500, // £5.00
				Currency: SupportedCurrencies["USD"],
			},
			expected:  &Money{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
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
		money1    *Money
		money2    *Money
		expected  *Money
		expectErr bool
	}{
		{
			name: "Subtract same currency",
			money1: &Money{
				Amount:   1000, // £10.00
				Currency: SupportedCurrencies["GBP"],
			},
			money2: &Money{
				Amount:   500, // £5.00
				Currency: SupportedCurrencies["GBP"],
			},
			expected: &Money{
				Amount:   500, // £15.00
				Currency: SupportedCurrencies["GBP"],
			},
			expectErr: false,
		},
		{
			name: "Subtract different currencies",
			money1: &Money{
				Amount:   1000, // £10.00
				Currency: SupportedCurrencies["GBP"],
			},
			money2: &Money{
				Amount:   500, // £5.00
				Currency: SupportedCurrencies["USD"],
			},
			expected:  &Money{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
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
		name           string
		money1         *Money
		exchangeRate   *Ratio
		targetCurrency Currency
		expected       *Money
		expectErr      bool
	}{
		{
			name: "Convert to different currency",
			money1: &Money{
				Amount:   1000, // $10.00
				Currency: SupportedCurrencies["GBP"],
			},
			exchangeRate:   exchangeRate,
			targetCurrency: SupportedCurrencies["USD"],
			expected: &Money{
				Amount:   600, // £15.00
				Currency: SupportedCurrencies["USD"],
			},
			expectErr: false,
		},
		{
			name: "Convert to same currencies",
			money1: &Money{
				Amount:   1000, // £10.00
				Currency: SupportedCurrencies["GBP"],
			},
			exchangeRate:   exchangeRate2,
			targetCurrency: SupportedCurrencies["GBP"],
			expected:       &Money{},
			expectErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			result, err := tt.money1.Convert(tt.targetCurrency.Code, tt.exchangeRate)

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
