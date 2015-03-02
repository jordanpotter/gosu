package db

import (
	"time"
)

type AccountConn interface {
	CreateAccount() (*Account, string, error)
	GetAccount(id, password string) (*Account, error)
}

type Account struct {
	Id      string
	Created time.Time
}
