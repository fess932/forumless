package models

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrUserNotFound Error = "user not found"
	ErrPostNotFound Error = "post not found"

	ErrWrongText Error = "wrong text in text field"

	ErrUserExists Error = "user same name already created"
	ErrUserEmpty  Error = "wrong user name"
)
