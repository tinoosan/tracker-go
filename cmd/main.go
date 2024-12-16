package main

import (
	"fmt"
	"time"
	"trackergo/internal/category"
	"trackergo/internal/transaction"
	"trackergo/internal/users"
)

func main() {

	var (
		userMap        = users.NewInMemoryStore()
		transactionMap = transaction.NewInMemoryStore()
		categoryMap    = category.NewInMemoryStore()

		username  = "Testuser1234"
		email     = "testuser@test.com"
		password  = "MyStrongPassword123!"
		createdAt = time.Now()
		amount    = 302.10
	)
	newUser, err := users.NewUser(username, email, password)
	if err != nil {
		fmt.Println(err)
		return
	}

  err = userMap.AddUser(newUser)
  if err != nil {
    fmt.Println(err)
    return
  }
	bills, err := category.NewCategory("bills", newUser.Id, false)
	if err != nil {
		fmt.Println(err)
    return
	}
  err = categoryMap.AddCategory(bills)
  if err != nil {
    fmt.Println(err)
    return
  }

  transaction, err := transaction.NewTransaction(bills.Id, newUser.Id, amount, createdAt)
  if err != nil {
    fmt.Println(err)
    return
  }

  err = transactionMap.AddTransaction(transaction)
  if err != nil {
    fmt.Println(err)
    return
  }
	result, err := transactionMap.ListTransactions(newUser.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
	fmt.Println(result[0].Id, result[0].CategoryID, result[0].Amount, result[0].CreatedAt)

}
