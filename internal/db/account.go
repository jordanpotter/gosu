package db

import (
	"time"
)

type AccountConn interface {
	CreateAccount(email, clientName, clientPassword string) error
	GetAccount(email string) (*Account, error)
}

type Account struct {
	Id      string
	Email   string
	Clients []Client
}

type Client struct {
	Name         string
	PasswordHash []byte
	Created      time.Time
}
