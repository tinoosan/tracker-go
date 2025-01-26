package valueobjects

type AccountDetails struct {
	Code Code
	Name string
	Type AccountType
}

func NewAccountDetails(code Code, name string, accountType AccountType) *AccountDetails {
	return &AccountDetails{
		Code: code,
		Name: name,
		Type: accountType,
	}
}
