package users

type Error struct {
	message string
}

var (
	ErrUsernameExists    = &Error{message: "Username already exists"}
	ErrUsernameNull      = &Error{message: "Username cannot be empty"}
	ErrUsernameInvalid   = &Error{message: "Username contains invalid characters. Only characters a-z, A-Z and spaces are allowed"}
	ErrUserIdNull        = &Error{message: "User ID cannot be empty"}
	ErrEmailExists       = &Error{message: "Email already exists"}
	ErrEmailNull         = &Error{message: "Email cannot be empty"}
	ErrEmailInvalid      = &Error{message: "Email contains invalid characters. Only characters a-z, A-Z and spaces are allowed"}
	rrPasswordNull       = &Error{message: "Password is required"}
	ErrPasswordTooShort  = &Error{message: "Password must be at least 8 characters long"}
	ErrPasswordNoLower   = &Error{message: "Password must contain at least one lowercase letter"}
	ErrPasswordNoUpper   = &Error{message: "Password must contain at least one uppercase letter"}
	ErrPasswordNoDigit   = &Error{message: "Password must contain at least one digit"}
	ErrPasswordNoSpecial = &Error{message: "Password must contain at least one special character (@$!%*?&#)"}
	ErrPasswordInvalid   = &Error{message: "Password contains invalid characters."}
	ErrUserNotCreated    = &Error{message: "User could not be created"}
	ErrUserNotFound      = &Error{message: "Category could not be found or does not exist"}
)

func (e *Error) Error() string {
	return e.message
}
