package main

import (
	"fmt"
	"os"
	"trackergo/internal/adapters/cli"
	"trackergo/internal/adapters/database/memory"
	"trackergo/internal/application"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)

type App struct {
	accountRepo   application.AccountRepository
	accService    *application.AccountService
	ledgerRepo    application.LedgerRepository
	ledgerService *application.LedgerService
}

func New(userID uuid.UUID) *App {
	accountRepo := memory.NewAccountMemoryStore()
	accService := application.NewAccountService(accountRepo)

	ledgerRepo := memory.NewLedgerMemoryStore()
	ledgerService := application.NewLedgerService(ledgerRepo, accService)

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

	utils.ShowMenu()
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
			return

		}
	}

}

func main() {
	userID := uuid.New()
	app := New(userID)
	app.run(userID)

}
