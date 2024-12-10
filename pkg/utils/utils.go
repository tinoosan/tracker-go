package utils

import "github.com/google/uuid"

func GenerateUUID() uuid.UUID {
	u := uuid.New()
	return u
}
