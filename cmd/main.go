package main

import (
	"fmt"
	"time"
	"trackergo/internal/category"
	"trackergo/internal/transaction"
)

func main() {
  c := category.NewInMemoryStore()
  c.CreateDefaultCategories()
  t := transaction.NewInMemoryStore()
  transaction := transaction.Transaction{}
  bills, err := category.NewCategory("bills")
  if err != nil {
    fmt.Println(err)
  }
  transaction.NewTransaction(time.Now(), bills, 1000.0)
  t.AddTransaction(transaction)
  result := t.ListTransactions()
  fmt.Println(result[0].Id, result[0].Category.Name, result[0].Amount, result[0].CreatedAt)


}
