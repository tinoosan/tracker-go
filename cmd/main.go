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

    userService = users.NewUserService(userMap)
    transactionService = transaction.NewTransactionService(transactionMap)

		username  = "Testuser1234"
		email     = "testuser@test.com"
		password  = "MyStrongPassword123!"
		createdAt = time.Now()
		amount    = 302.10
	)

  newUser, err := userService.CreateUser(username, email, password)
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

  _, err = transactionService.CreateTransaction(newUser.Id, bills.Id, amount, createdAt)
  if err != nil {
    fmt.Println(err)
    return
  }

	result, err := transactionService.GetAllTransaction(newUser.Id)
  if err != nil {
    fmt.Println(err)
    return
  }
	fmt.Println(result[0].Id, result[0].CategoryID, result[0].Amount, result[0].CreatedAt)

}
