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
	CreatedAt     time.Time
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
	fmt.Println("Getting transactions...")
	for k, v := range transactions {
		result = append(result, fmt.Sprintf("\n\n ID: %v\n Category: %v\n Amount: %v\n Created at: %v\n\n", k, v.Category.Name, v.Amount, v.CreatedAt.Format("2006-01-02 15:04:05")))
	}
	return result
}

func CreateTransaction(createdAt time.Time, category *category.Category, amount float64) error {
	if createdAt.String() == "" {
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

	t := Transaction{Id: utils.GenerateUUID(), CreatedAt: createdAt, Category: category, Amount: amount}
	fmt.Println("Creating transaction with id: ", t.Id)
	addTransaction(t, transactions)

	return nil
}
