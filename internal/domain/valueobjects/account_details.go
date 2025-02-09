package valueobjects

import "fmt"

type AccountDetails struct {
	Code Code
	Name string
	Type AccountType
}

func NewAccountDetails(code Code, name string, accountType AccountType) (*AccountDetails, error) {

  if name == "" {
    return &AccountDetails{}, fmt.Errorf("name cannot be empty")
  }

	return &AccountDetails{
		Code: code,
		Name: name,
		Type: accountType,
	},nil
}


