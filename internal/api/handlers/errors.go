package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

type Error struct {
	message string
}

var (
	ErrUserIDRequired        = &Error{message: "User ID is required"}
	ErrUserIDInvalid         = &Error{message: "UserID format is invalid"}
	ErrTransactionIDRequired = &Error{message: "Transaction ID format is invalid"}
	ErrTransactionIDInvalid  = &Error{message: "Transaction ID format is invalid"}
	ErrCategoryIDRequired    = &Error{message: "Category ID is required"}
	ErrCreatingCategory      = &Error{message: "Category could not be created"}
	ErrFetchingCategory      = &Error{message: "Category could not be fetched"}
	ErrUpdatingCategory      = &Error{message: "Category could not be updated"}
	ErrDeletingCategory      = &Error{message: "Category could not be deleted"}
	ErrCreatingTransaction   = &Error{message: "Transaction could not be created"}
	ErrFetchingTransaction   = &Error{message: "Transaction could not be fetched"}
	ErrUpdatingTransaction   = &Error{message: "Transaction could not be updated"}
	ErrDeletingTransaction   = &Error{message: "Transaction could not be deleted"}
)

func (e *Error) Error() string {
	return e.message
}

func WriteJSONError(w http.ResponseWriter, statusCode int, message string, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Error:   message,
		Details: details,
	}
	json.NewEncoder(w).Encode(response)
}
