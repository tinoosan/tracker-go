package transaction

import (
	"testing"
	"time"
  "trackergo/internal/category"
)

var (
	categoryName = "bills"
	bills    = category.CreateCategory(categoryName)
	date     = time.Now()
	amount   = 302.10
)

func TestCreateTransaction(t *testing.T) {
	err := CreateTransaction(date, bills, amount)
	if err != nil {
		t.Error(err)
	}
}
