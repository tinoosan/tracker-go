package main

import (
	"fmt"
	"log"
	"net/http"
	api "trackergo/internal/api/handlers"
	"trackergo/internal/category"
	"trackergo/internal/transaction"
	"trackergo/internal/users"
	"trackergo/middleware"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	categoryRepo := category.NewInMemoryStore()
	categoryService := category.NewCategoryService(categoryRepo)
	categoryHandler := api.NewCategoryHandler(categoryService)

	userRepo := users.NewInMemoryStore()
	userService := users.NewUserService(userRepo, categoryService)
	userHandler := api.NewUserHandler(userService)

	transactionRepo := transaction.NewInMemoryStore()
	transactionService := transaction.NewTransactionService(transactionRepo, categoryService)
	transactionHandler := api.NewTransactionHandler(transactionService)

	router := mux.NewRouter()
	handler := enableCORS(router)

	router.HandleFunc("/api/v1/login", userHandler.Login).Methods("POST")
	router.HandleFunc("/api/v1/logout", userHandler.Logout).Methods("POST")
	router.HandleFunc("/api/v1/register", userHandler.CreateUser).Methods("POST")

	router.Handle("/api/v1/users", middleware.RequireAuth(http.HandlerFunc(userHandler.GetUser))).Methods("GET")
	router.Handle("/api/v1/users", middleware.RequireAuth(http.HandlerFunc(userHandler.UpdateUser))).Methods("PATCH")
	router.Handle("/api/v1/users", middleware.RequireAuth(http.HandlerFunc(userHandler.DeleteUser))).Methods("DELETE")

	router.Handle("/api/v1/users/categories", middleware.RequireAuth(http.HandlerFunc(categoryHandler.CreateCategory))).Methods("POST")
	router.Handle("/api/v1/users/categories/{id}", middleware.RequireAuth(http.HandlerFunc(categoryHandler.GetCategoryByID))).Methods("GET")
	router.Handle("/api/v1/users/categories/{id}", middleware.RequireAuth(http.HandlerFunc(categoryHandler.UpdateCategory))).Methods("PATCH")
	router.Handle("/api/v1/users/categories/{id}/reactivate", middleware.RequireAuth(http.HandlerFunc(categoryHandler.ReactivateCategory))).Methods("POST")
	router.Handle("/api/v1/users/categories/{id}", middleware.RequireAuth(http.HandlerFunc(categoryHandler.DeleteCategory))).Methods("DELETE")
	router.Handle("/api/v1/users/categories", middleware.RequireAuth(http.HandlerFunc(categoryHandler.GetAllCategories))).Methods("GET")

	router.Handle("/api/v1/users/transactions", middleware.RequireAuth(http.HandlerFunc(transactionHandler.CreateTransaction))).Methods("POST")
	router.Handle("/api/v1/users/transactions/{id}", middleware.RequireAuth(http.HandlerFunc(transactionHandler.GetTransactionByID))).Methods("GET")
	router.Handle("/api/v1/users/transactions/{id}", middleware.RequireAuth(http.HandlerFunc(transactionHandler.UpdateTransaction))).Methods("PATCH")
	router.Handle("/api/v1/users/transactions/{id}", middleware.RequireAuth(http.HandlerFunc(transactionHandler.DeleteTransaction))).Methods("DELETE")
	router.Handle("/api/v1/users/transactions", middleware.RequireAuth(http.HandlerFunc(transactionHandler.GetUserTransactions))).Methods("GET")

	fmt.Println("Server is running on port 8080..")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
