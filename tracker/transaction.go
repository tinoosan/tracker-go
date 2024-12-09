package tracker

import (
	"time"
)

type Transaction struct {
	Id       int
	Date     time.Time
	Category Category
	Amount   float64
}

