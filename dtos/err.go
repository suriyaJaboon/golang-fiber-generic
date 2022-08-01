package dtos

import "strconv"

type (
	Error struct {
		Code int
		Opt  string
		Err  error
	}
)

func (e *Error) Error() string { return e.Opt + " " + strconv.Itoa(e.Code) + ": " + e.Err.Error() }
