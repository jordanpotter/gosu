package postgres

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedAccount struct {
	ID      int       `db:"id"`
	Email   string    `db:"email"`
	Created time.Time `db:"created"`
}

func (sa *storedAccount) toAccount() *db.Account {
	return &db.Account{
		ID:      sa.ID,
		Email:   sa.Email,
		Created: sa.Created,
	}
}

func (c *conn) CreateAccount(email, deviceName, devicePassword string) (*db.Account, error) {
	tx, err := c.Beginx()
	if err != nil {
		return nil, err
	}

	sa := new(storedAccount)
	// TODO: use ON CONFLICT in Postgres 9.5 to do nothing if account already exists
	insertAccount := "INSERT INTO accounts (email, created) VALUES ($1, $2) RETURNING *"
	err = tx.Get(sa, insertAccount, email, time.Now())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	devicePasswordHash, err := password.ComputeHash(devicePassword)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	insertDevice := "INSERT INTO devices (account_id, name, password_hash, created) VALUES ($1, $2, $3, $4)"
	_, err = tx.Exec(insertDevice, sa.ID, deviceName, devicePasswordHash, time.Now())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	return sa.toAccount(), err
}

func (c *conn) GetAccount(id int) (*db.Account, error) {
	sa := new(storedAccount)
	query := "SELECT * FROM accounts WHERE id=$1 LIMIT 1"
	err := c.Get(sa, query, id)
	return sa.toAccount(), err
}

func (c *conn) GetAccountByEmail(email string) (*db.Account, error) {
	sa := new(storedAccount)
	query := "SELECT * FROM accounts WHERE email=$1 LIMIT 1"
	err := c.Get(sa, query, email)
	return sa.toAccount(), err
}
