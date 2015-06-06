package db

type Conn struct {
	// Accounts AccountsConn
	// Rooms    RoomsConn
	Closer
}

type Closer interface {
	Close() error
}
