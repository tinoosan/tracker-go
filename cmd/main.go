package main

import (
	"fmt"
	"log"
	"net/http"
	"trackergo/internal/api/handlers"
	"trackergo/internal/category"
	"trackergo/internal/transaction"
	"trackergo/internal/users"
	"trackergo/middleware"

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

  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/index.html")
  })

	router.HandleFunc("/login", userHandler.Login).Methods("POST")
	router.HandleFunc("/logout", userHandler.Logout).Methods("POST")

	router.HandleFunc("/register", userHandler.CreateUser).Methods("POST")
	router.Handle("/users", middleware.RequireAuth(http.HandlerFunc(userHandler.GetUser))).Methods("GET")
	router.Handle("/users", middleware.RequireAuth(http.HandlerFunc(userHandler.UpdateUser))).Methods("PATCH")
	router.Handle("/users", middleware.RequireAuth(http.HandlerFunc(userHandler.DeleteUser))).Methods("DELETE")

	router.Handle("/users/categories", middleware.RequireAuth(http.HandlerFunc(categoryHandler.CreateCategory))).Methods("POST")
	router.Handle("/users/categories/{id}", middleware.RequireAuth(http.HandlerFunc(categoryHandler.GetCategoryByID))).Methods("GET")
	router.Handle("/users/categories/{id}", middleware.RequireAuth(http.HandlerFunc(categoryHandler.UpdateCategory))).Methods("PATCH")
	router.Handle("/users/categories/{id}", middleware.RequireAuth(http.HandlerFunc(categoryHandler.DeleteCategory))).Methods("DELETE")
	router.Handle("/users/categories", middleware.RequireAuth(http.HandlerFunc(categoryHandler.GetAllCategories))).Methods("GET")

	router.Handle("/users/transactions", middleware.RequireAuth(http.HandlerFunc(transactionHandler.CreateTransaction))).Methods("POST")
	router.Handle("/users/transactions/{id}", middleware.RequireAuth(http.HandlerFunc(transactionHandler.GetTransactionByID))).Methods("GET")
	router.Handle("/users/transactions/{id}", middleware.RequireAuth(http.HandlerFunc(transactionHandler.UpdateTransaction))).Methods("PATCH")
	router.Handle("/users/transactions/{id}", middleware.RequireAuth(http.HandlerFunc(transactionHandler.DeleteTransaction))).Methods("DELETE")
	router.Handle("/users/transactions", middleware.RequireAuth(http.HandlerFunc(transactionHandler.GetUserTransactions))).Methods("GET")

	fmt.Println("Server is running on port 8080..")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
