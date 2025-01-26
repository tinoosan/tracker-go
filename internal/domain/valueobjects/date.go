package valueobjects

import (
	"time"
)

type Date struct {
	Value string
}

func NewDate(value time.Time) *Date {
	return &Date{Value: value.Format("2006-01-02")}
}
