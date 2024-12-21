package main

import (
	"fmt"
	"log"
	"net/http"
	"trackergo/internal/api/handlers"
	"trackergo/internal/category"
	"trackergo/internal/transaction"
	"trackergo/internal/users"

	"github.com/gorilla/mux"
)

func main() {

	userRepo := users.NewInMemoryStore()
	userService := users.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

  categoryRepo := category.NewInMemoryStore()
  categoryService := category.NewCategoryService(categoryRepo)
  categoryHandler := api.NewCategoryHandler(categoryService)

  categoryService.CreateDefaultCategories()

  transactionRepo := transaction.NewInMemoryStore()
  transactionService := transaction.NewTransactionService(transactionRepo)
  transactionHandler := api.NewTransactionHandler(transactionService)
  
	router := mux.NewRouter()

	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	router.HandleFunc("/users/{userId}/categories", categoryHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/users/{userId}/categories/{id}", categoryHandler.GetCategoryByID).Methods("GET")
	router.HandleFunc("/users/{userId}/categories/{id}", categoryHandler.UpdateCategory).Methods("PATCH")
	router.HandleFunc("/users/{userId}/categories/{id}", categoryHandler.DeleteCategory).Methods("DELETE")
  router.HandleFunc("/users/{userId}/categories", categoryHandler.GetAllCategories).Methods("GET")

  router.HandleFunc("/users/{userId}/transactions", transactionHandler.CreateTransaction).Methods("POST")
	router.HandleFunc("/users/{userId}/transactions/{id}", transactionHandler.GetTransactionByID).Methods("GET")
	router.HandleFunc("/users/{userId}/transactions/{id}", transactionHandler.UpdateTransaction).Methods("PATCH")
	router.HandleFunc("/users/{userId}/transactions/{id}", transactionHandler.DeleteTransaction).Methods("DELETE")
  router.HandleFunc("/users/{userId}/transactions", transactionHandler.GetUserTransactions).Methods("GET")



	fmt.Println("Server is running on port 8080..")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
