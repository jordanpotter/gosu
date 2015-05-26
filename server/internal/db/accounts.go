package db

import (
	"time"
)

type AccountsConn interface {
	Create(email, deviceName, devicePassword string) (*Account, error)
	Get(id string) (*Account, error)
	GetByEmail(email string) (*Account, error)
	Delete(id string) error
}

type Account struct {
	ID      string    `json:"id"`
	Email   string    `json:"email"`
	Devices []Device  `json:"devices"`
	Created time.Time `json:"created"`
}

type Device struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	PasswordHash []byte    `json:"-"`
	Created      time.Time `json:"created"`
}
