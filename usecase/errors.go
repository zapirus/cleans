package usecase

import (
	"errors"

	"github.com/jackc/pgx"
)

const (
	ErrorNotFound Error = iota + 1
	OnAutorization
	ErrorLimit
)

type Error uint8

func (e Error) Error() string {
	switch {
	case errors.Is(e, ErrorNotFound):
		return "not found"
	case errors.Is(e, OnAutorization):
		return "on autorization"
	case errors.Is(e, ErrorLimit):
		return "limit error"
	}
	return "unknown error"
}

func wrapError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrorNotFound
	}
	return err
}
