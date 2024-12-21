package api

import (
	"encoding/json"
	"net/http"
	"trackergo/internal/users"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service users.UserService
}

func NewUserHandler(service users.UserService) *UserHandler {
	return &UserHandler{Service: service}
}


// POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newUser, err := h.Service.CreateUser(userRequest.Username, userRequest.Email, userRequest.Password)
	if err != nil {
    WriteJSONError(w, http.StatusInternalServerError, "User could not be created", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// GET /users/{id}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	userID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
	}

	user, err := h.Service.GetUserByID(userID)
	if err != nil {
    WriteJSONError(w, http.StatusNotFound, "User Not Found", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// PATCH /users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
		return
	}

	var userRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
    WriteJSONError(w, http.StatusInternalServerError, "User could not be updated", err.Error())
		return
	}

	updatedUser, err := h.Service.UpdateUser(userID, userRequest.Username, userRequest.Email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)

}

// DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	userID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
	}

	err = h.Service.DeleteUser(userID)
	if err != nil {
    WriteJSONError(w, http.StatusInternalServerError, "User could not be deleted", err.Error())
    return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
