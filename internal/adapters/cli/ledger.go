package cli

import (
	"fmt"
	"os"
	"strings"
	"trackergo/internal/application"
	"trackergo/internal/domain/ledger"
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

	debitTxn, creditTxn, err := service.CreateTransaction(debitAccount, creditAccount, userID, amount, ledger.GBP, description)
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
