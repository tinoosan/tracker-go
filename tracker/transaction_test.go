package tracker

import (
	"testing"
	"time"
)

var (
	category = "bills"
	bills    = CreateCategory(category)
	date     = time.Now()
	amount   = 302.10
)

func TestCreateTransaction(t *testing.T) {
	err := CreateTransaction(date, bills, amount)
	if err != nil {
		t.Error(err)
	}
}
