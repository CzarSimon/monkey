package object

import "fmt"

// Error Object for wrapping an encountered error
type Error struct {
	Message string
}

func (err *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (err *Error) Inspect() string {
	return "ERROR: " + err.Message
}

// NewError Creates an error based on a message and returns a referece to it
func NewError(message string) *Error {
	return &Error{
		Message: message,
	}
}

// NewErrorf Formats an error messages and creates an Error object
func NewErrorf(format string, a ...interface{}) *Error {
	return NewError(fmt.Sprintf(format, a...))
}
