package users

import "github.com/google/uuid"

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
	return nil, &Error{}
}

func (s *userService) GetUserByID(userId uuid.UUID) (*User, error) {
	return nil, &Error{}
}

func (s *userService) UpdateUser(userId uuid.UUID, username, email string) (*User, error) {
	return nil, &Error{}
}

func (s *userService) DeleteUser(userId uuid.UUID) error {
	return nil
}
