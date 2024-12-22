package api

import (
	"encoding/json"
	"net/http"
	"trackergo/internal/category"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	Service *category.CategoryService
}

func NewCategoryHandler(service *category.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: service}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var userRequest struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		WriteJSONError(w, http.StatusBadRequest, "Invalid request payoad", err.Error())
		return
	}

	newCategory, err := h.Service.CreateCategory(userID, userRequest.Name)
	if err != nil {
		WriteJSONError(w, http.StatusInternalServerError, ErrCreatingCategory.message, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}

func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingCategory.message, ErrCategoryIDRequired.message)
		return
	}
	categoryId, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingCategory.message, err.Error())
		return
	}

	category, err := h.Service.GetCategoryById(categoryId, userID)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingCategory.message, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id, ok := vars["id"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrUpdatingCategory.message, ErrCategoryIDRequired.message)
		return
	}
	categoryId, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrUpdatingCategory.message, err.Error())
		return
	}

	var userRequest struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		WriteJSONError(w, http.StatusBadRequest, "Invalid request payoad", err.Error())
		return
	}

	category, err := h.Service.UpdateCategory(categoryId, userID, userRequest.Name)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrUpdatingCategory.message, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  userID, ok := r.Context().Value("userId").(uuid.UUID)
  if !ok {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
  }

  id, ok := vars["id"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, ErrCategoryIDRequired.message)
		return
	}
	categoryId, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, err.Error())
		return
	}

	err = h.Service.DeleteCategory(categoryId, userID)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
  userID, ok := r.Context().Value("userId").(uuid.UUID)
  if !ok {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
  }

	userCategories, err := h.Service.GetAllCategories(userID)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userCategories)
}
