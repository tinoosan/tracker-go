package api

import (
	"encoding/json"
	"net/http"
	"time"
	"trackergo/internal/users"
	"trackergo/middleware"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
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
  utils.WriteJSONResponse(w, http.StatusCreated, newUser)
}

// POST /login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		WriteJSONError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	user, err := h.Service.AuthenticateUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		WriteJSONError(w, http.StatusUnauthorized, "Invalid email or password", err.Error())
		return
	}

	sessionID := middleware.CreateSession(user.Id)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
  utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Login successful"})
}

// POST /logout
func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		middleware.DeleteSession(cookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfuly"})
}

// GET /users/
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.Service.GetUserByID(userID)
	if err != nil {
		WriteJSONError(w, http.StatusNotFound, "User Not Found", err.Error())
		return
	}
  utils.WriteJSONResponse(w, http.StatusOK, user)
}

// PATCH /users
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
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
  if err != nil {
    WriteJSONError(w, http.StatusInternalServerError, "User could not be updated", err.Error())
    return
  }
  utils.WriteJSONResponse(w, http.StatusOK, updatedUser)

}

// DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
  err := h.Service.DeleteUser(userID)
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, "User could not be deleted", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
