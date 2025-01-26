package ledger

import (
	"fmt"
)

type Money struct {
	Amount   int
	Currency Currency
}

type Currency struct {
  Code string
  SubUnit int
  Symbol string
}

type Percentage struct {
  Value float64
}

type Tax struct {
  Rate Percentage
  Amount Money
}


var SupportedCurrencies = map[string]Currency{
  "GBP": Currency{Code:"GBP", SubUnit: 100, Symbol: "£"},
  "USD": Currency{Code:"USD", SubUnit: 100, Symbol: "$"},
  "EUR": Currency{Code:"EUR", SubUnit: 100, Symbol: "€"},
  "JPY": Currency{Code:"JPY", SubUnit: 1, Symbol: "¥"},
}

func NewTax(rate Percentage, base Money) ( *Tax, error) {
  taxAmount := base.Amount * int(rate.Value*100)
  return &Tax{
    Rate: rate,
    Amount: Money{Amount: taxAmount, Currency: base.Currency},
  }, nil
}

func NewMoney(amount float64, currency string) (*Money, error) {
	if !isSupportedCurrency(currency) {
		return &Money{}, fmt.Errorf("unsupported currency: %s", currency)
	}
	if amount < 0 {
		return &Money{}, fmt.Errorf("amount cannot be negative")
	}
	return &Money{
		Amount:   int(amount * 100),
		Currency: SupportedCurrencies[currency],
	}, nil
}

// To convert amount subunit to unit depending on currency
func (m *Money) GetAmount() float64 {
    return float64(m.Amount)/float64(m.Currency.SubUnit)
  }

func (m *Money) Add(other *Money) (*Money, error) {
	if m.Currency != other.Currency {
		return nil, fmt.Errorf("cannot add different currencies: %s and %s",
			m.Currency.Code, other.Currency.Code)
	}
	return &Money{
		Amount:   m.Amount + other.Amount,
		Currency: m.Currency,
	}, nil
}

func (m *Money) Subtract(other *Money) (*Money, error) {
	if m.Currency != other.Currency {
		return nil, fmt.Errorf("cannot subtract different currencies: %s and %s",
			m.Currency.Code, other.Currency.Code)
	}
	return &Money{
		Amount:   m.Amount - other.Amount,
		Currency: m.Currency,
	}, nil
}

// Convert the Money struct into a different currency using an exchange
// rate

func (m *Money) Convert(targetCurrency string, exchangeRate float64) (*Money, error) {
	if exchangeRate <= 0 {
		return &Money{}, fmt.Errorf("invalid exchange rate: %.2f", exchangeRate)
	}

  if targetCurrency == m.Currency.Code {
    return &Money{}, fmt.Errorf("cannot convert to same current currency")
  }

	if !isSupportedCurrency(targetCurrency) {
		return &Money{}, fmt.Errorf("unsupported target currency: %s", targetCurrency)
	}
	return &Money{
		Amount:   (m.Amount * int(exchangeRate*100))/100,
		Currency: SupportedCurrencies[targetCurrency],
	}, nil
}

func (m *Money) Format() string {
	return fmt.Sprintf("%.2f %s", m.GetAmount(), m.Currency.Code)
}

func isSupportedCurrency(currency string) bool {
	_, ok := SupportedCurrencies[currency]
	return ok
}
