package cli

import (
	"fmt"
	"os"
	"strings"
	"trackergo/internal/accounts"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)

func AccountsMenu(service *accounts.Service, userID uuid.UUID) {
	var choice int

	fmt.Println("Choose an option: ")
	fmt.Println("1. Create Account")
	fmt.Println("2. View Account")
	fmt.Println("3. Update Account")
	fmt.Println("4. Chart Of Accounts")
	fmt.Println("5. Deactivate Account")
	fmt.Println("6. Exit")
	fmt.Println("7. Main Menu")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		createAccount(service, userID)
	case 2:
		viewAccount(service, userID)
	case 3:
		updateAccount(service, userID)
	case 4:
		viewChartOfAccounts(service, userID)
	case 5:
		deleteAccount(service, userID)
	case 6:
		fmt.Println("Exiting application")
		os.Exit(0)
	case 7:
		utils.ShowMenu()
	default:
		fmt.Println("Invalid choice. Please try again")
	}

}

func createAccount(service *accounts.Service, userID uuid.UUID) {
  defer utils.ShowMenu()
	name, err := utils.GetInputString("Enter account name: ")
	if err != nil {
		fmt.Println("Error: ", err)
	}

  name = strings.Trim(name, "\n")
  name = strings.ToUpper(name)
	accountType, err := utils.GetInputString("Enter account type (ASSET, LIABILITY, EQUITY, EXPENSE, REVENUE): ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	accountType = strings.Trim(accountType, "\n")
  accountType = strings.ToUpper("accountType")
	newAccount, err := service.CreateAccount(userID, name, accounts.Type(accountType))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Account '%s - %v' has been created\n", newAccount.Name, newAccount.Code)

}

func viewAccount(service *accounts.Service, userID uuid.UUID) {
  defer utils.ShowMenu()
	var name string

	fmt.Print("Enter account name: ")
	fmt.Scan(&name)
  name = strings.ToUpper(name)
	account, err := service.GetAccountByName(name, userID)
	if err != nil || account == nil {
		fmt.Printf(err.Error(), name)
		return
	}

	fmt.Printf("Account Code: %v\nAccount Name: %s\nType: %s\nCurrent Balance:%.2f\n",
		account.Code, account.Name, account.Type, account.CurrentBalance())
}

func updateAccount(service *accounts.Service, userID uuid.UUID) {
  defer utils.ShowMenu()
	var code int
	var name string

	fmt.Print("Enter account code: ")
	fmt.Scan(&code)

	err := service.UpdateAccount(accounts.Code(code), userID, name)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Account has been updated successfully")
}

func deleteAccount(service *accounts.Service, userID uuid.UUID) {
  defer utils.ShowMenu()
	var code int

	fmt.Print("Enter account code: ")
	fmt.Scan(&code)

	err := service.DeleteAccount(accounts.Code(code), userID)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func viewChartOfAccounts(service *accounts.Service, userID uuid.UUID) {
  defer utils.ShowMenu()
	accounts, err := service.GetChartOfAccounts(userID)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("=================================================================")
	fmt.Println("                      CHART OF ACCOUNTS                          ")
	fmt.Println("=================================================================")
	fmt.Printf("%-20s %-24s %-10s\n","Account Code", "Account Name", "Account Type")
	fmt.Println("-----------------------------------------------------------------")

	for _, account := range accounts {
		fmt.Printf("%-20v %-24s %-10s\n", account.Code, account.Name, account.Type)
	fmt.Println("-----------------------------------------------------------------")
	}
}