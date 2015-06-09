package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate/migrate"

	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type conn struct {
	*sqlx.DB
}

func New(addrs []config.PostgresNode, config *config.Postgres) (db.Conn, error) {
	fmt.Println("TODO: support multiple postgres endpoints")
	postgres, err := getDBWithAddr(addrs[0], config)
	if err != nil {
		return nil, err
	}

	err = performMigrations(addrs[0], config, "./conf/db/migrations")
	if err != nil {
		return nil, err
	}

	return &conn{postgres}, nil
}

func getDBWithAddr(addr config.PostgresNode, config *config.Postgres) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		addr.IP.String(), addr.DBPort, config.Name, config.Username, config.Password, config.SSLMode)
	return sqlx.Connect("postgres", connStr)
}

func performMigrations(addr config.PostgresNode, config *config.Postgres, migrationsPath string) error {
	connStr := fmt.Sprintf("postgres://%s:%d/%s?user=%s&password=%s&sslmode=%s",
		addr.IP.String(), addr.DBPort, config.Name, config.Username, config.Password, config.SSLMode)
	allErrors, ok := migrate.UpSync(connStr, migrationsPath)
	if !ok {
		return allErrors[0]
	}
	return nil
}
