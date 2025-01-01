package cli

import (
	"fmt"
	"trackergo/internal/accounts"
	"trackergo/internal/ledger"
)

func TAccountHeader(account *accounts.Account) {
	fmt.Println("==========================================================================================================")
	fmt.Printf("                                                   %s                                               \n", account.Name)
	fmt.Println("==========================================================================================================")
	fmt.Printf("%-20s %-20s %-24s %-15s %-15s %-15s\n", "Date","Txn No.", "Description", "Debit", "Credit", "Balance")
	fmt.Println("----------------------------------------------------------------------------------------------------------")

}

func TAccountItem(entries []*ledger.Entry) {
	for _, entry := range entries {
		if entry.EntryType == ledger.Debit {
			fmt.Printf("%-20s %-20v %-24s %-15.2f %-15.2f %-15.2f\n", entry.CreatedAt.Format("2006-01-02"), entry.ID, entry.Description, entry.Amount, 0.00, entry.GetBalance())
			return
		}
		fmt.Printf("%-20s %-20v %-24s %-15.2f %-15.2f %-15.2f\n", entry.CreatedAt.Format("2006-01-02"), entry.ID, entry.Description, 0.00, entry.Amount, entry.GetBalance())
	}

}

func TAccountBalance(entries []*ledger.Entry) {
	var totalBalance float64
	for _, entry := range entries {
		totalBalance += entry.GetBalance()
	}
	fmt.Println("==========================================================================================================")
	fmt.Printf("%-100s %.2f\n", "", totalBalance)
}
