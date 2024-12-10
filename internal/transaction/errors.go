package transaction

type Error struct {
	message string
}
var (
	ErrDateNull                = &Error{message: "Date is required"}
	ErrAmountNull              = &Error{message: "Amount is required"}
	ErrAmountNotPositive       = &Error{message: "Amount must be positive"}
	ErrTransactionCategoryNull = &Error{message: "Category is required"}
)

func (e *Error) Error() string {
	return e.message
}
