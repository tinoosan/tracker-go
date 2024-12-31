package cli

import (
	"fmt"
	"os"
	"trackergo/internal/accounts"

	"github.com/google/uuid"
)

func AccountsMenu(service *accounts.Service, userID uuid.UUID) {
	var choice int

	fmt.Println("Choose an option: ")
	fmt.Println("1. Create Account")
	fmt.Println("2. View Account")
	fmt.Println("3. Update Account")
	fmt.Println("4. Deactivate Account")
	fmt.Println("5. Exit")
	fmt.Println("6. Main Menu")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		createAccount(*service, userID)
	case 2:
		viewAccount(*service, userID)
	case 3:
		updateAccount(*service, userID)
	case 4:
		deleteAccount(*service, userID)
	case 5:
		fmt.Println("Exiting application")
		os.Exit(0)
	case 6:
		return
	default:
		fmt.Println("Invalid choice. Please try again")
	}

}

func createAccount(service accounts.Service, userID uuid.UUID) {
	var name string
	var accountType string

	fmt.Print("Enter account name: ")
	fmt.Scan(&name)

	fmt.Print("Enter account type (ASSET, LIABILITY, EQUITY, EXPENSE, REVENUE): ")
	fmt.Scan(&accountType)

	account, err := service.CreateAccount(userID, name, accounts.Type(accountType))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Account created successfully with ID:", account.Id)

}

func viewAccount(service accounts.Service, userID uuid.UUID) {
	var name string

	fmt.Print("Enter account name: ")
	fmt.Scan(&name)

	account, err := service.GetAccountByName(name, userID)
	if err != nil || account == nil {
		fmt.Printf(err.Error(), name)
		return
	}

  fmt.Printf("Account ID: %v\nName: %s\nType: %s\nCurrent Balance:%.2f\n",
		account.Id, account.Name, account.Type, account.CurrentBalance())
}

func updateAccount(service accounts.Service, userID uuid.UUID) {
	var accountID, name string

	fmt.Print("Enter account ID: ")
	fmt.Scan(&accountID)

	id, err := uuid.Parse(accountID)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = service.UpdateAccount(id, userID, name)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Account has been updated successfully")
}

func deleteAccount(service accounts.Service, userID uuid.UUID) {
	var accountID string

	fmt.Print("Enter account ID: ")
	fmt.Scan(&accountID)

	id, err := uuid.Parse(accountID)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = service.DeleteAccount(id, userID)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
