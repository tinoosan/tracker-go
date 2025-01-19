package users

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Username string
	Email    string
	Password string
}

func NewUser(username, email, password string) *User {

	return &User{
		Id:       uuid.New(),
		Username: username,
		Email:    email,
		Password: password,
	}
}
