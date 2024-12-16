package transaction

type Error struct {
	message string
}

var (
	ErrDateNull                    = &Error{message: "Date is required"}
	ErrAmountNull                  = &Error{message: "Amount is required"}
	ErrAmountNotPositive           = &Error{message: "Amount must be positive"}
	ErrTransactionCategoryNull     = &Error{message: "Category is required"}
	ErrTransactionNull             = &Error{message: "Transaction has not been created before adding to store"}
	ErrTransactionNotFound         = &Error{message: "Transaction with ID '%v' could not be found"}
	ErrTransactionWithUserNotFound = &Error{message: "Transaction with ID '%s' with user ID '%s' could not be found"}
)

func (e *Error) Error() string {
	return e.message
}
