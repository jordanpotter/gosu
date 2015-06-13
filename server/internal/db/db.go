package db

import "io"

type Conn interface {
	AccountsConn
	DevicesConn
	io.Closer
}

type AccountsConn interface {
	CreateAccount(email string) (*Account, error)
	GetAccount(id int) (*Account, error)
	GetAccountByEmail(email string) (*Account, error)
}

type DevicesConn interface {
	CreateDevice(accountID int, deviceName, devicePassword string) (*Device, error)
	GetDevices(accountID int) ([]Device, error)
}
