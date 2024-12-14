package users

import (
	"regexp"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Username string
	Email    string
	Password string
}

type UserRepository interface {
	AddUser(u *User) error
	GetUserByID(userId uuid.UUID) (*User, error)
	UpdateUserByID(userID uuid.UUID, username, email string) (*User, error)
}

type InMemoryStore struct {
	Users            map[uuid.UUID]*User
	UserIDToUsername map[string]uuid.UUID
	UserIDToEmail    map[string]uuid.UUID
}

type Error struct {
	message string
}

var (
	ErrUsernameExists    = &Error{message: "Username already exists"}
	ErrUsernameNull      = &Error{message: "Username cannot be empty"}
	ErrUsernameInvalid   = &Error{message: "Username contains invalid characters. Only characters a-z, A-Z and spaces are allowed"}
	ErrUserIdNull        = &Error{message: "User ID cannot be empty"}
	ErrEmailExists       = &Error{message: "Email already exists"}
	ErrEmailNull         = &Error{message: "Email cannot be empty"}
	ErrEmailInvalid      = &Error{message: "Email contains invalid characters. Only characters a-z, A-Z and spaces are allowed"}
	rrPasswordNull       = &Error{message: "Password is required"}
	ErrPasswordTooShort  = &Error{message: "Password must be at least 8 characters long"}
	ErrPasswordNoLower   = &Error{message: "Password must contain at least one lowercase letter"}
	ErrPasswordNoUpper   = &Error{message: "Password must contain at least one uppercase letter"}
	ErrPasswordNoDigit   = &Error{message: "Password must contain at least one digit"}
	ErrPasswordNoSpecial = &Error{message: "Password must contain at least one special character (@$!%*?&#)"}
	ErrPasswordInvalid   = &Error{message: "Password contains invalid characters."}
	ErrUserNotCreated    = &Error{message: "User could not be created"}
	ErrUserNotFound      = &Error{message: "Category could not be found or does not exist"}
)

var _ UserRepository = &InMemoryStore{}

func (e *Error) Error() string {
	return e.message
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{Users: make(map[uuid.UUID]*User),
		UserIDToUsername: make(map[string]uuid.UUID),
		UserIDToEmail:    make(map[string]uuid.UUID)}
}

func NewUser(username, email, password string) (*User, error) {
	if username == "" {
		return nil, ErrUsernameNull
	}
	if !isUsernameValid(username) {
		return nil, ErrUsernameInvalid
	}

	if email == "" {
		return nil, ErrEmailNull
	}

	if !isEmailValid(email) {
		return nil, ErrEmailInvalid
	}

	_, err := isPasswordValid(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:       utils.GenerateUUID(),
		Username: username,
		Email:    email,
		Password: password,
	}, nil
}

func (s *InMemoryStore) AddUser(user *User) error {
	if user == nil {
		return ErrUserNotCreated
	}
	_, ok := s.UserIDToEmail[user.Email]
	if ok {
		return ErrEmailExists
	}

	_, ok = s.UserIDToUsername[user.Username]
	if ok {
		return ErrUsernameExists
	}
	s.UserIDToEmail[user.Email] = user.Id
	s.UserIDToUsername[user.Username] = user.Id
	s.Users[user.Id] = user
	return nil

}

func (s *InMemoryStore) GetUserByID(userId uuid.UUID) (*User, error) {
	if userId.String() == "" {
		return nil, ErrUserIdNull
	}
	user, exists := s.userIDExists(userId)
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *InMemoryStore) UpdateUserByID(userId uuid.UUID, username, email string) (*User, error) {
	user, exists := s.userIDExists(userId)
	if !exists {
		return nil, ErrUserNotFound
	}
 if username != user.Username {
    user.Username = username
  }
 if email != user.Email {
    user.Email = email
  }
	return user, nil
}

func isEmailValid(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(email) {
		return false
	}
	return true
}

func isUsernameValid(username string) bool {
	pattern := `^[a-zA-Z][a-zA-Z0-9_]{3,15}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(username) {
		return false
	}
	return true
}

func (s *InMemoryStore) userIDExists(userId uuid.UUID) (*User, bool) {
	user, ok := s.Users[userId]
	if !ok {
		return nil, false
	}
	return user, true
}

func isPasswordValid(password string) (bool, error) {
	if len(password) < 8 {
		return false, ErrPasswordTooShort
	}

	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false, ErrPasswordNoLower
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false, ErrPasswordNoUpper
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		return false, ErrPasswordNoDigit
	}

	if !regexp.MustCompile(`[@$!%*?&#]`).MatchString(password) {
		return false, ErrPasswordNoSpecial
	}

	if regexp.MustCompile(`[^A-Za-z\d@$!%*?&#]`).MatchString(password) {
		return false, ErrPasswordInvalid
	}

	return true, nil

}
