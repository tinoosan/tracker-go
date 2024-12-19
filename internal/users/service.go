package users

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(username, email, password string) (*User, error)
	GetUserByID(userId uuid.UUID) (*User, error)
	UpdateUser(userId uuid.UUID, username, email string) (*User, error)
	DeleteUser(userId uuid.UUID) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(username, email, password string) (*User, error) {
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

	newUser := NewUser(username, email, password)
	if newUser == nil {
		return nil, ErrUserNotCreated
	}

	err = s.repo.AddUser(newUser)
	if err != nil {
		return nil, err
	}
  fmt.Printf("User has been created with id %v\n", newUser.Id)
	return newUser, nil
}

func (s *userService) GetUserByID(userId uuid.UUID) (*User, error) {
	if userId.String() == "" {
		return nil, ErrUserIdNull
	}

  user, err := s.repo.GetUserByID(userId)
  if err != nil {
    return nil, err
  }

	return user, nil
}

func (s *userService) UpdateUser(userId uuid.UUID, username, email string) (*User, error) {
  updatedUser, err := s.repo.UpdateUserByID(userId, username, email)
  if err != nil {
    return nil, err
  }
	return updatedUser, nil
}

func (s *userService) DeleteUser(userId uuid.UUID) error {
  err := s.repo.DeleteUserByID(userId)
  if err != nil {
    return err
  }
	return nil
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
