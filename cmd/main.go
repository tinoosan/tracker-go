package main

import (
	"fmt"
	"os"
	"trackergo/cmd/cli"
	"trackergo/internal/accounts"
	"trackergo/internal/ledger"

	"github.com/google/uuid"
)

type App struct {
	accountRepo   *accounts.InMemoryStore
	accService    *accounts.Service
	ledgerRepo    ledger.LedgerRepository
	ledgerService *ledger.Service
}

func New(userID uuid.UUID) *App {
	accountRepo := accounts.NewInMemoryStore()
	accService := accounts.NewService(accountRepo)

	ledgerRepo := ledger.NewInMemoryStore()
	ledgerService := ledger.NewService(ledgerRepo, *accService)

	if _, ok := accountRepo.UserAccounts[userID]; !ok {
		err := accService.CreateDefaultAccounts(userID)
		if err != nil {
			fmt.Printf("Failed to create default accounts: %v\n", err)
			os.Exit(1)
		}

	}
	return &App{
		accountRepo:   accountRepo,
		accService:    accService,
		ledgerRepo:    ledgerRepo,
		ledgerService: ledgerService,
	}
}

func (a *App) run(userID uuid.UUID) {
menuOptions := []string{"Accounts", "Transactions", "Exit"}
	fmt.Println("Choose an option:")
	for i, option := range menuOptions {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	for {
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			cli.AccountsMenu(a.accService, userID)
		case 2:
			cli.TransactionsMenu(a.ledgerService, userID)
		case 3:
			fmt.Println("Exiting application...")
			return
		default:
			fmt.Println("Invalid option. Please try again")

		}
	}

 }

func main() {
	userID := uuid.New()
	app := New(userID)
	app.run(userID)

}
