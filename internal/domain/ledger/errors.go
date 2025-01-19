package ledger

type Error struct {
	message string
}

var (
	ErrDateNull                = &Error{message: "Date is required"}
	ErrAmountNull              = &Error{message: "Amount is required"}
	ErrAmountNotPositive       = &Error{message: "Amount must be positive"}
	ErrEntryCategoryNull = &Error{message: "Category is required"}
	ErrEntryNull         = &Error{message: "Entry has not been created before adding to store"}
	ErrEntryNotFound     = &Error{message: "Entry with ID '%v' could not be found"}
	ErrUserEntrysNotFound = &Error{message: "No entries found"}
)

func (e *Error) Error() string {
	return e.message
}
