package valueobjects

import "fmt"

type Money struct {
	Amount   int
	Currency Currency
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
	return float64(m.Amount) / float64(m.Currency.SubUnit)
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

func (m *Money) Convert(targetCurrency string, exchangeRate *Ratio) (*Money, error) {
	if targetCurrency == m.Currency.Code {
		return &Money{}, fmt.Errorf("cannot convert to same current currency")
	}

	if !isSupportedCurrency(targetCurrency) {
		return &Money{}, fmt.Errorf("unsupported target currency: %s", targetCurrency)
	}

  convertedAmount := exchangeRate.Apply(m.Amount)
	return &Money{
		Amount:   convertedAmount,
		Currency: SupportedCurrencies[targetCurrency],
	}, nil
}

func (m *Money) Format() string {
	return fmt.Sprintf("%.2f %s", m.GetAmount(), m.Currency.Code)
}
