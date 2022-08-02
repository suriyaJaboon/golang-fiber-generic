package services

import (
	"fg/dtos"
	"fg/x"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	SERVICE    = x.SERVICE
	SERVER     = x.SERVER
	NOTFOUND   = x.NOTFOUND
	FORMATTING = x.FORMATTING
)

var (
	ErrInvalid         = errInvalid()
	ErrInvalidFormat   = errInvalidFormat()
	ErrInvalidNotFound = errInvalidNotFound()
)

type Error = dtos.Error

func ErrServer(err error) error {
	return &Error{Opt: SERVER, Err: err}
}

func ErrByID(err error) error {
	if err == mongo.ErrNoDocuments {
		return ErrInvalid
	}

	return ErrServer(err)
}

func errInvalid() error {
	return &Error{Opt: SERVICE, Err: os.ErrInvalid}
}

func errInvalidFormat() error {
	return &Error{Opt: FORMATTING, Err: os.ErrInvalid}
}

func errInvalidNotFound() error {
	return &Error{Opt: NOTFOUND, Err: os.ErrInvalid}
}
