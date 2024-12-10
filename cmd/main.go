package main

import (
	"fmt"
	"time"
	"trackergo/internal/category"
	"trackergo/internal/transaction"
)

func main() {
  c := category.NewCategories()
  c.CreateDefaultCategories()
  t := transaction.NewTransactionsMap()
  transaction := transaction.Transaction{}
  transaction.NewTransaction(time.Now(), c.Store["bills"], 1000.0)
  t.AddTransaction(transaction)

  fmt.Println(t.ListTransactions())


}
