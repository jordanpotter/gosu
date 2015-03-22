package db

type Conn interface {
	AccountsConn
	RoomsConn
	Close()
}
