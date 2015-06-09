package db

import "time"

type Account struct {
	ID      int
	Email   string
	Created time.Time
}

type Device struct {
	ID           int
	Name         string
	PasswordHash []byte
	Created      time.Time
	LastLogin    time.Time
}
