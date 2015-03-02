package db

type Conn interface {
	AccountConn
	Close()
}
