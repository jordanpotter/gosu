package db

import (
	"fmt"
)

const (
	minAccountNameLength = 16
	maxAccountNameLength = 32
)

type Account struct {
	Id           string `json:"-", bson:"-"`
	Name         string `json:"name", bson:"name"`
	PasswordHash []byte `json:"passwordHash", bson:"passwordHash"`
}

func CheckAccountName(name string) error {
	if len(name) < minAccountNameLength || len(name) > maxAccountNameLength {
		return fmt.Errorf("db: invalid name length %d", len(name))
	}
	return nil
}
