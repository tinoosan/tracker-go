package transaction

import (
	"testing"
	"time"
  "trackergo/internal/category"
)

var (
	categoryName = "bills"
  transactions = NewTransactionsMap()
	bills    = category.CreateCategory(categoryName)
	createdAt     = time.Now()
	amount   = 302.10
  transaction = Transaction{}
)

func TestCreateTransaction(t *testing.T) {
  transaction.NewTransaction(createdAt, bills, amount)
	err := transactions.AddTransaction(transaction)
	if err != nil {
		t.Error(err)
	}
}
