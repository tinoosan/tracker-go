package cli

import (
	"fmt"
	"os"
	"strconv"
	"trackergo/internal/accounts"
	"trackergo/internal/ledger"

	"github.com/olekukonko/tablewriter"
)

func TAccountHeader(account *accounts.Account) {
	fmt.Println("====================================================================================================")
	fmt.Printf("                                                   %s                                               \n", account.Name)
	fmt.Println("====================================================================================================")
}

func TAccountTable(entries []*ledger.Entry) {

	data := [][]string{}
	var totalBalance float64

	for _, entry := range entries {
		var debitAmount, creditAmount string
		date := entry.CreatedAt.Format("2006-01-02")
		txnID := entry.ID.String()

		switch entry.EntryType {
		case ledger.Debit:
			debitAmount = strconv.FormatFloat(entry.Amount, 'f', 2, 64)
		case ledger.Credit:
			creditAmount = strconv.FormatFloat(entry.Amount, 'f', 2, 64)

		}

		balance := strconv.FormatFloat(entry.GetBalance(), 'f', 2, 64)
		row := []string{date, txnID, entry.Description, debitAmount, creditAmount, balance}
		data = append(data, row)
	}

	for _, entry := range entries {
		totalBalance += entry.GetBalance()
	}
	lastRow := []string{"", "", "", "", "", strconv.FormatFloat(totalBalance, 'f', 2, 64)}
	data = append(data, lastRow)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Txn ID", "Description", "Debit", "Credit", "Balance"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()

}
