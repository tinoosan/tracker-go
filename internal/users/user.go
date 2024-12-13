package users

import (
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID
}

func NewUser() *User {
  user := &User{
    Id: utils.GenerateUUID(),
  }

  return user
}
