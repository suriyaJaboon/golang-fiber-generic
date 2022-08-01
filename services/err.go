package services

import (
	"fg/dtos"
	"fg/x"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
)

const (
	SERVER   = x.SERVER
	NOTFOUND = x.NOTFOUND
)

var (
	ErrInvalid = errInvalid()
)

type Error = dtos.Error

func ErrServer(err error) error {
	return &Error{Code: http.StatusInternalServerError, Opt: SERVER, Err: err}
}

func ErrByID(err error) error {
	if err == mongo.ErrNoDocuments {
		return ErrInvalid
	}

	return ErrServer(err)
}

func errInvalid() error {
	return &Error{Code: http.StatusNotFound, Opt: NOTFOUND, Err: os.ErrInvalid}
}
