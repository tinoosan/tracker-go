package valueobjects


type Tax struct {
	Rate   Percentage
	Amount Money
}

func NewTax(rate Percentage, base Money) (*Tax, error) {
	taxAmount := rate.Apply(base.Amount)
	return &Tax{
		Rate:   rate,
		Amount: Money{Amount: taxAmount, Currency: base.Currency},
	}, nil
}

