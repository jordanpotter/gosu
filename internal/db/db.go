package db

type Conn interface {
	Close()
	CreateAccount(name, password string) (*Account, error)
	GetAccount(name, password string) (*Account, error)
}
