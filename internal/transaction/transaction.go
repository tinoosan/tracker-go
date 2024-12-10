package transaction

import (
	"fmt"
	"time"
	"trackergo/internal/category"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)

type Transaction struct {
	Id       uuid.UUID
	Date     time.Time
	Category *category.Category
	Amount   float64
	created  bool
}

type Error struct {
	message string
}

var (
	transactions = make(map[uuid.UUID]*Transaction)
)

var (
	ErrDateNull                = &Error{message: "Date is required"}
	ErrAmountNull              = &Error{message: "Amount is required"}
	ErrAmountNotPositive       = &Error{message: "Amount must be positive"}
	ErrTransactionCategoryNull = &Error{message: "Category is required"}
)

func (e *Error) Error() string {
	return e.message
}

func addTransaction(t Transaction, transactions map[uuid.UUID]*Transaction) {
	for {
		_, ok := transactions[t.Id]
		if !ok {
			break
		}
		t.Id = utils.GenerateUUID()
	}

	transactions[t.Id] = &t
}

func ListTransactions() []string {
	var result []string
	for k, v := range transactions {
		result = append(result, fmt.Sprintf("ID: %v, Category: %v, Amount: %v, Date: %v", k, v.Category.Name, v.Amount, v.Date.String()))
	}
	return result
}

func CreateTransaction(date time.Time, category *category.Category, amount float64) error {
	if date.String() == "" {
		return ErrDateNull
	}
	if category == nil {
		return ErrTransactionCategoryNull
	}
	if amount == 0.0 {
		return ErrAmountNull
	}
	if amount < 0.0 {
		return ErrAmountNotPositive
	}
	t := Transaction{Id: utils.GenerateUUID(), Date: date, Category: category, Amount: amount}
	addTransaction(t, transactions)

	return nil
}
