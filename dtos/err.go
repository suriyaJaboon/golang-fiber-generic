package dtos

import "strconv"

type (
	Error struct {
		Opt string
		Err error
	}
	APIError struct {
		Code int
		Opt  string
		Err  error
	}
)

func (e *Error) Error() string    { return e.Opt + ": " + e.Err.Error() }
func (e *APIError) Error() string { return e.Opt + " " + strconv.Itoa(e.Code) + ": " + e.Err.Error() }
