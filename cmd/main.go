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

  for _, v := range c.Store {
    transaction.CreateTransaction(time.Now(), v, 1000.0)
  }

  fmt.Println(transaction.ListTransactions())


}
