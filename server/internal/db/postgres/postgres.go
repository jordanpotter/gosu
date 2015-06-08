package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func New(addrs []config.PostgresNode, config *config.Postgres) (*db.Conn, error) {
	postgresConn, err := getConnWithAddr(addrs[0], config)
	if err != nil {
		return nil, err
	}

	conn := &db.Conn{
		Closer: postgresConn,
	}
	return conn, nil

	// accountsConn, err := accounts.New(session, config)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// roomsConn, err := rooms.New(session, config)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// conn := &db.Conn{
	// 	Accounts: accountsConn,
	// 	Rooms:    roomsConn,
	// 	Closer:   session,
	// }
	// return conn, nil
}

func getConnWithAddr(addr config.PostgresNode, config *config.Postgres) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		addr.IP.String(), addr.DBPort, config.Username, config.Password, config.Name, config.SSLMode)
	return sqlx.Connect("postgres", connStr)
}
