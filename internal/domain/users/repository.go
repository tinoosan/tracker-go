package users

import (
	"errors"
	"github.com/google/uuid"
)



type UserRepository interface {
	AddUser(u *User) error
	GetUserByID(userId uuid.UUID) (*User, error)
  GetUserByEmail(email string) (*User, error)
	UpdateUserByID(userID uuid.UUID, username, email string) (*User, error)
	DeleteUserByID(userID uuid.UUID) error
}

type InMemoryStore struct {
	Users            map[uuid.UUID]*User
	UserIDToUsername map[string]uuid.UUID
	UserIDToEmail    map[string]uuid.UUID
}

var _ UserRepository = &InMemoryStore{}
var SystemUser = &User{
	Username: "system",
	Email:    "system@tracker.com",
	Password: "",
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{Users: make(map[uuid.UUID]*User),
		UserIDToUsername: make(map[string]uuid.UUID),
		UserIDToEmail:    make(map[string]uuid.UUID)}
}


func (s *InMemoryStore) AddUser(user *User) error {
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
		user, exists := s.userIDExists(userId)
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *InMemoryStore) GetUserByEmail(email string) (*User, error) {
  userId, exists := s.UserIDToEmail[email]
  if !exists {
    return nil, errors.New("Email does not exist")
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
	if username != user.Username && username != ""{
		user.Username = username
	}
	if email != user.Email && email != "" {
		user.Email = email
	}
	return user, nil
}

func (s *InMemoryStore) DeleteUserByID(userId uuid.UUID) error {
	_, exists := s.userIDExists(userId)
	if !exists {
		return ErrUserNotFound
	}
	delete(s.Users, userId)
	return nil
}


func (s *InMemoryStore) userIDExists(userId uuid.UUID) (*User, bool) {
	user, ok := s.Users[userId]
	if !ok {
		return nil, false
	}
	return user, true
}


