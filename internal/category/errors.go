package category

type Error struct {
	message string
}

var (
	ErrCategoryExists      = &Error{message: "Category already exists"}
	ErrCategoryNull        = &Error{message: "Name cannot be empty"}
	ErrCategoryInvalid     = &Error{message: "Name contains invalid characters. Only characters a-z, A-Z and spaces are allowed"}
	ErrCategoryNotAdded    = &Error{message: "Category could not be added"}
	ErrCategoryNotFound    = &Error{message: "Category could not be found or does not exist"}
	ErrUserHasNoCategories = &Error{message: "No categories found for user with id %v"}
  ErrCategoryHasNoUser = &Error{message: "User is nil in category. Make sure the user is assigned to a category"}
  ErrCategoryIdNull = &Error{message: "Category Id is required"}
)

func (e *Error) Error() string {
	return e.message
}
