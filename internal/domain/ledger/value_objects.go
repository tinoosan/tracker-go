package ledger

import (
	"fmt"
)

type Money struct {
	Amount   int
	Currency Currency
}

type Currency string

const (
  GBP Currency = "GBP"
  EUR Currency = "EUR"
  USD Currency = "USD"
)
var SupportedCurrencies = map[Currency]bool{
	GBP: true,
	EUR: true,
	USD: true,
}

var SubUnits = map[Currency]int{
  GBP: 100,
  EUR: 100,
  USD: 100,
}

func NewMoney(amount float64, currency Currency) (*Money, error) {
	if !isSupportedCurrency(currency) {
		return &Money{}, fmt.Errorf("unsupported currency: %s", currency)
	}
	if amount < 0 {
		return &Money{}, fmt.Errorf("amount cannot be negative")
	}
	return &Money{
		Amount:   int(amount * 100),
		Currency: currency,
	}, nil
}

// To convert amount subunit to unit depending on currency
func (m *Money) GetAmount() float64 {
  subUnit, _ := SubUnits[m.Currency]
    return float64(m.Amount)/float64(subUnit)
  }

func (m *Money) Add(other *Money) (*Money, error) {
	if m.Currency != other.Currency {
		return nil, fmt.Errorf("cannot add different currencies: %s and %s",
			m.Currency, other.Currency)
	}
	return &Money{
		Amount:   m.Amount + other.Amount,
		Currency: m.Currency,
	}, nil
}

func (m *Money) Subtract(other *Money) (*Money, error) {
	if m.Currency != other.Currency {
		return nil, fmt.Errorf("cannot subtract different currencies: %s and %s",
			m.Currency, other.Currency)
	}
	return &Money{
		Amount:   m.Amount - other.Amount,
		Currency: m.Currency,
	}, nil
}

// Convert the Money struct into a different currency using an exchange
// rate

func (m *Money) Convert(targetCurrency Currency, exchangeRate float64) (*Money, error) {
	if exchangeRate <= 0 {
		return &Money{}, fmt.Errorf("invalid exchange rate: %.2f", exchangeRate)
	}

  if targetCurrency == m.Currency {
    return &Money{}, fmt.Errorf("cannot convert to same current currency")
  }

	if !isSupportedCurrency(targetCurrency) {
		return &Money{}, fmt.Errorf("unsupported target currency: %s", targetCurrency)
	}
	return &Money{
		Amount:   (m.Amount * int(exchangeRate*100))/100,
		Currency: targetCurrency,
	}, nil
}

func (m *Money) Format() string {
	return fmt.Sprintf("%.2f %s", float64(m.Amount)/100, m.Currency)
}

func isSupportedCurrency(currency Currency) bool {
	_, ok := SupportedCurrencies[currency]
	return ok
}
