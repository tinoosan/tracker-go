package cli

import (
	"fmt"
	"os"
	"trackergo/internal/ledger"

	"github.com/google/uuid"
)

func TransactionsMenu(service *ledger.Service, userID uuid.UUID) {
	var choice int

	fmt.Print("Choose an option: ")
	fmt.Println("1. Create Transaction")
	fmt.Println("2. View Transaction")
	fmt.Println("3. Update Transaction")
	fmt.Println("4. Reverse Transaction")
	fmt.Println("5. Exit")
	fmt.Println("6. Main Menu")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		createTransaction(*service, userID)
	case 2:
		viewTAccount(*service, userID)
	case 3:
		updateTransaction(*service, userID)
	case 4:
		reverseTransaction(*service, userID)
	case 5:
		fmt.Println("Exiting application")
		os.Exit(0)
	case 6:
		return
	default:
		fmt.Println("Invalid choice. Please try again")
	}

}

func createTransaction(service ledger.Service, userID uuid.UUID) {
	var debitAccount, creditAccount, description string
	var amount float64

	fmt.Print("Enter an account to debit: ")
	fmt.Scan(&debitAccount)

	fmt.Print("Enter an account to credit: ")
	fmt.Scan(&creditAccount)

	fmt.Print("Enter an amount: ")
	fmt.Scan(&amount)

	fmt.Print("Enter a description: ")
	fmt.Scan(&description)

	debitTxn, creditTxn, err := service.CreateTransaction(debitAccount, creditAccount, userID, amount, description)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Transactions processed: %v, %v ", debitTxn, creditTxn)

}

func viewTAccount(service ledger.Service, userID uuid.UUID) {
	var name string
	var totalBalance float64
	fmt.Print("Enter the account name: ")
	fmt.Scan(&name)

	result, account, err := service.GetTAccount(name, userID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("----------------------------------------------")
	fmt.Printf("                      %s                       \n", account.Name)
	fmt.Println("----------------------------------------------")
	fmt.Printf("Date           | Debit   | Credit   | Balance\n")
	for _, entry := range result {
		if entry.EntryType == ledger.Debit {

			fmt.Printf("%s       | %.2f      |%.2f    | %.2f\n", entry.CreatedAt.Format("2006-01-02"), entry.Amount, 0.00, entry.GetBalance())
			return
		}
		fmt.Printf("%s     | %.2f    | %.2f   | %.2f\n", entry.CreatedAt.Format("2006-01-02"), 0.00, entry.Amount, entry.GetBalance())
	}

	for _, entry := range result {
		totalBalance += entry.GetBalance()
	}
  fmt.Println("----------------------------------------------")

	fmt.Printf("                              Total | %.2f\n", totalBalance)

}

func updateTransaction(service ledger.Service, userID uuid.UUID) {
	return
}

func reverseTransaction(service ledger.Service, userID uuid.UUID) {
	return
}
