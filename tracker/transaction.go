package tracker

import (
	"time"
  "github.com/google/uuid"
)

type Transaction struct {
	Id       uuid.UUID
	Date     time.Time
	Category Category
	Amount   float64
}

