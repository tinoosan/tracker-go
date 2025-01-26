package valueobjects


type Currency struct {
	Code    string
	SubUnit int
	Symbol  string
}


var SupportedCurrencies = map[string]Currency{
	"GBP": Currency{Code: "GBP", SubUnit: 100, Symbol: "£"},
	"USD": Currency{Code: "USD", SubUnit: 100, Symbol: "$"},
	"EUR": Currency{Code: "EUR", SubUnit: 100, Symbol: "€"},
	"JPY": Currency{Code: "JPY", SubUnit: 1, Symbol: "¥"},
}


func isSupportedCurrency(currency string) bool {
	_, ok := SupportedCurrencies[currency]
	return ok
}
