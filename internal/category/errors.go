package category

type Error struct {
	message string
}

var (
	ErrCategoryExists   = &Error{message: "Category already exists"}
	ErrCategoryNull     = &Error{message: "Name cannot be empty"}
	ErrCategoryInvalid  = &Error{message: "Name contains invalid characters. Only characters a-z, A-Z and spaces are allowed"}
	ErrCategoryNotAdded = &Error{message: "Category could not be added"}
	ErrCategoryNotFound = &Error{message: "Category could not be found or does not exist"}
)

func (e *Error) Error() string {
	return e.message
}
