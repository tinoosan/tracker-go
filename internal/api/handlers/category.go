package api

import (
	"encoding/json"
	"net/http"
	"trackergo/internal/category"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	Service category.CategoryService
}

func NewCategoryHandler(service category.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: service}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrCreatingCategory.message, ErrUserIDRequired.message)
		return
	}
	userID, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrCreatingCategory.message, ErrUserIDInvalid.message)
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
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingCategory.message, ErrUserIDRequired.message)
	}
	userId, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingCategory.message, err.Error())
		return
	}

	id, ok = vars["id"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingCategory.message, ErrCategoryIDRequired.message)
		return
	}
	categoryId, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrFetchingCategory.message, err.Error())
		return
	}

	category, err := h.Service.GetCategoryById(categoryId, userId)
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
	id, ok := vars["userId"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrUpdatingCategory.message, ErrUserIDRequired.message)
	}
	userId, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrUpdatingCategory.message, err.Error())
		return
	}

	id, ok = vars["id"]
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

	category, err := h.Service.UpdateCategory(categoryId, userId, userRequest.Name)
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
	id, ok := vars["userId"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, ErrUserIDRequired.message)
	}
	userId, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, err.Error())
		return
	}

	id, ok = vars["id"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, ErrCategoryIDRequired.message)
		return
	}
	categoryId, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, err.Error())
		return
	}

	err = h.Service.DeleteCategory(categoryId, userId)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
 vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, ErrUserIDRequired.message)
	}
	userId, err := uuid.Parse(id)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, err.Error())
		return
	}

  userCategories, err := h.Service.GetAllCategories(userId)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, ErrDeletingCategory.message, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(userCategories)
}

