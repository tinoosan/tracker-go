package users

import (
	"errors"
	"fmt"
	"regexp"
	"trackergo/internal/category"

	"github.com/google/uuid"
)

type UserService struct {
	repo            UserRepository
  categoryService *category.CategoryService
}

func NewUserService(repo UserRepository, categoryService *category.CategoryService) *UserService {
	return &UserService{repo: repo, categoryService: categoryService}
}

func (s *UserService) CreateUser(username, email, password string) (*User, error) {
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

  err = s.categoryService.CreateDefaultCategories(newUser.Id)
  if err != nil {
    return nil, err
  }
	fmt.Printf("User has been created with id %v\n", newUser.Id)
	return newUser, nil
}

func (s *UserService) GetUserByID(userId uuid.UUID) (*User, error) {
	if userId.String() == "" {
		return nil, ErrUserIdNull
	}

	user, err := s.repo.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(userId uuid.UUID, username, email string) (*User, error) {
	updatedUser, err := s.repo.UpdateUserByID(userId, username, email)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (s *UserService) DeleteUser(userId uuid.UUID) error {
	err := s.repo.DeleteUserByID(userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) AuthenticateUser(email, password string) (*User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil || user.Password != password {
		return nil, errors.New("Invalid email or password")
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
