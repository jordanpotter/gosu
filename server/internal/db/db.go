package db

import "io"

type Conn interface {
	AccountsConn
	// DevicesConn
	io.Closer
}

type AccountsConn interface {
	CreateAccount(email, deviceName, devicePassword string) (*Account, error)
	GetAccount(id int) (*Account, error)
	GetAccountByEmail(email string) (*Account, error)
}

// type DevicesConn interface {
// 	GetDevices(accountID int) ([]Device, error)
// }
