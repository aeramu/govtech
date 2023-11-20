package errorhelper

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	code int
	err  error
}

func (e *Error) Error() string {
	return e.err.Error()
}

func New(msg string) error {
	return errors.New(msg)
}

func Wrap(err error, msg string) error {
	if err == nil {
		return New(msg)
	}
	return fmt.Errorf("%w: %w", New(msg), err)
}

func NewWithCode(msg string, code int) error {
	return &Error{
		code: code,
		err:  errors.New(msg),
	}
}

func WrapWithCode(err error, msg string, code int) error {
	if err == nil {
		return NewWithCode(msg, code)
	}
	return &Error{
		code: code,
		err:  fmt.Errorf("%w: %w", NewWithCode(msg, code), err),
	}
}

func GetCode(err error) int {
	errCode, ok := err.(*Error)
	if !ok {
		return http.StatusOK
	}
	return errCode.code
}
