package cli

import (
	"fmt"
	"os"
	"strings"
	"trackergo/internal/application"
	vo "trackergo/internal/domain/valueobjects"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)

func TransactionsMenu(service *application.LedgerService, userID uuid.UUID) {
	var choice int

	fmt.Println("Choose an option: ")
	fmt.Println("1. Create Entry")
	fmt.Println("2. View T-Account")
	fmt.Println("3. Reverse Entry")
	fmt.Println("4. Exit")
	fmt.Println("5. Main Menu")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		createTransaction(*service, userID)
	case 2:
		viewTAccount(*service, userID)
	case 3:
		reverseTransaction(*service, userID)
	case 4:
		fmt.Println("Exiting application")
		os.Exit(0)
	case 5:
		utils.ShowMenu()
	default:
		fmt.Println("Invalid choice. Please try again")
	}

}

func createTransaction(service application.LedgerService, userID uuid.UUID) {
	defer utils.ShowMenu()

	  var currencies []vo.Currency
	for _, currency := range vo.SupportedCurrencies {
    currencies = append(currencies, currency)
	}


  fmt.Println("Available Currencies:")
  for i, currency := range currencies {
    fmt.Printf("%d. %s\n", i+1, currency.Code)
  }

  var choice int
  fmt.Print("Enter the number corresponding to the currency: ")
  _, err := fmt.Scan(&choice)
  if err != nil || choice < 1 || choice > len(vo.SupportedCurrencies) {
    fmt.Println("Invalid selection. Please enter a valid number.")
    return
  }

  selectedCurrency := currencies[choice-1].Code
  fmt.Println("You selected:", selectedCurrency)

	debitAccount, err := utils.GetInputString("Enter an account to debit: ")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	debitAccount = strings.Trim(debitAccount, "\n")
	debitAccount = strings.ToUpper(debitAccount)

	creditAccount, err := utils.GetInputString("Enter an account to credit: ")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	creditAccount = strings.Trim(creditAccount, "\n")
	creditAccount = strings.ToUpper(creditAccount)

	amount, err := utils.GetInputFloat("Enter an amount: ")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	description, err := utils.GetInputString("Enter a description: ")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	description = strings.Trim(description, "\n")

	debitTxn, creditTxn, err := service.CreateTransaction(debitAccount, creditAccount, userID, amount, "GBP", description)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Transactions processed: %v, %v ", debitTxn, creditTxn)

}

func viewTAccount(service application.LedgerService, userID uuid.UUID) {
	defer utils.ShowMenu()
	var name string
	fmt.Print("Enter the account name: ")
	fmt.Scan(&name)

	name = strings.ToUpper(name)

	entries, account, err := service.GetTAccount(name, userID)
	if err != nil {
		fmt.Println(err)
	}

	TAccountHeader(account)
	TAccountTable(entries)

}

func reverseTransaction(service application.LedgerService, userID uuid.UUID) {

}
