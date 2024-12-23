package api

import (
	"encoding/json"
	"net/http"
	"time"
	"trackergo/internal/transaction"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	Service *transaction.TransactionService
}

func NewTransactionHandler(service *transaction.TransactionService) *TransactionHandler {
	return &TransactionHandler{Service: service}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
		return
	}

	var userRequest struct {
		CategoryId uuid.UUID `json:"categoryId"`
		Amount     float64   `json:"amount"`
    Description string `json:"description"`
		CreatedAt  time.Time `json:"createdAt"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		WriteJSONError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	newTransaction, err := h.Service.CreateTransaction(userID, userRequest.CategoryId, userRequest.Amount, userRequest.Description, userRequest.CreatedAt)
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, ErrCreatingTransaction.message, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTransaction)
}

func (h *TransactionHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingTransaction.message, ErrTransactionIDRequired.message)
		return
	}
	transactionId, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingTransaction.message, ErrTransactionIDInvalid.message)
		return
	}

	transaction, err := h.Service.GetTransactionById(transactionId, userID)
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, ErrFetchingTransaction.message, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingTransaction.message, ErrTransactionIDRequired.message)
		return
	}
	transactionID, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingTransaction.message, ErrTransactionIDInvalid.message)
		return
	}

	var userRequest struct {
		CategoryId uuid.UUID `json:"categoryId"`
		Amount     *float64  `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		WriteJSONError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	updatedTransaction, err := h.Service.UpdateTransaction(transactionID, userID, userRequest.CategoryId, userRequest.Amount)
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, ErrUpdatingTransaction.message, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedTransaction)
}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingTransaction.message, ErrTransactionIDRequired.message)
		return
	}
	transactionID, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingTransaction.message, ErrTransactionIDInvalid.message)
		return
	}
	err = h.Service.DeleteTransaction(transactionID, userID)
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, ErrDeletingTransaction.message, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *TransactionHandler) GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
		return
	}

	transactions, err := h.Service.GetAllTransactions(userID)
	if err != nil {
    WriteJSONError(w, http.StatusInternalServerError,err.Error(), "")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}
