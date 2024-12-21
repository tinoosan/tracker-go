package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	sessionStore = make(map[string]Session)
	mutex        = &sync.Mutex{}
)

type Session struct {
	UserID    uuid.UUID
	ExpiresAt time.Time
}

func CreateSession(userID uuid.UUID) string {
	mutex.Lock()
	defer mutex.Unlock()

	sessionID := uuid.NewString()
	sessionStore[sessionID] = Session{
		UserID:    userID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

  fmt.Printf("Created session with id %v", sessionID)

	return sessionID

}

func GetSession(sessionID string) (Session, error) {
	mutex.Lock()
	defer mutex.Unlock()

	session, exists := sessionStore[sessionID]
	if !exists || session.ExpiresAt.Before(time.Now()) {
		return Session{}, errors.New("Session expired or not found")
	}
	return session, nil
}

func DeleteSession(sessionID string) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(sessionStore, sessionID)
}

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    cookie, err := r.Cookie("session_id")
    if err != nil {
      http.Error(w, "Unauthorized: No session found", http.StatusUnauthorized)
      return
    }

    sessionID := cookie.Value
    session, err := GetSession(sessionID)
    if err != nil {
      http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
      return
    }

    ctx := context.WithValue(r.Context(), "userId", session.UserID)

    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

