package utils

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

func GenerateUUID() uuid.UUID {
	u := uuid.New()
	return u
}

func IsAuthorizedAndValid(w http.ResponseWriter, r *http.Request) (uuid.UUID, bool) {
	userID, ok := r.Context().Value("userId").(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
		return userID, false
	}
  return userID, true
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func WriteHTMLResponse(w http.ResponseWriter, statusCode int, templateName string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

  err = tmpl.Execute(w, data)
  if err != nil {
    http.Error(w, "Failed to render template", http.StatusInternalServerError)
    return
  }

}
