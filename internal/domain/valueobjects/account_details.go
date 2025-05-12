package valueobjects

import "fmt"

type AccountDetails struct {
	Code Code
	Name string
	Type AccountType
}

// NewAccountDetails creates and validates a new AccountDetails instance.
// Returns an error if the name is empty.
func NewAccountDetails(code Code, name string, accountType AccountType) (*AccountDetails, error) {

  if name == "" {
    return nil, fmt.Errorf("name cannot be empty")
  }

	return &AccountDetails{
		Code: code,
		Name: name,
		Type: accountType,
	},nil
}


