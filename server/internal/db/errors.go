package db

import (
	"errors"
)

var (
	NotFoundError  = errors.New("not found")
	DuplicateError = errors.New("duplicate")
	NotEmptyError  = errors.New("not empty")
)
