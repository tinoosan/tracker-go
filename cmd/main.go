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
    categoryService = category.NewCategoryService(categoryMap)
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
  err = categoryService.CreateDefaultCategories()
  newCategory, err := categoryService.CreateCategory(newUser.Id, "new category")
	
  _, err = transactionService.CreateTransaction(newUser.Id, newCategory.Id, amount, createdAt)
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
