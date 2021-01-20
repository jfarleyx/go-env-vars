package env

import (
	"errors"
	"fmt"
)

// ErrorType defines the type of error the env module returns.
type ErrorType string

const (
	// NotFound ErrorType indicates the requested env variable was not found.
	NotFound ErrorType = "NotFound"
	// EmptyValue ErrorType indicates the requested env variable's value was an empty string.
	EmptyValue ErrorType = "EmptyValue"
	// InvalidType ErrorType indicates the requested env variable's value was not of the requested data type.
	InvalidType ErrorType = "InvalidType"
)

// Error wraps a Go error and defines an ErrorType.
type Error struct {
	// Type is the ErrorType assigned to this Error.
	Type ErrorType
	// Msg is the message assigned to this Error.
	Err error
}

// Error returns an error message.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Err)
}

// NewError creates and returns a new env.Error{}.
func NewError(etype ErrorType, envVarName string, msg string) *Error {
	var m string = ""

	switch etype {
	case NotFound:
		m = fmt.Sprintf("env variable not found: %s;", envVarName)
	case EmptyValue:
		m = fmt.Sprintf("env variable is empty: %s;", envVarName)
	case InvalidType:
		m = fmt.Sprintf("invalid data type for env var: %s;", envVarName)
	}

	if msg != "" {
		m += " " + msg
	}

	return &Error{
		Type: etype,
		Err:  errors.New(m),
	}
}
