package postgres

import (
	"time"

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

func (c *conn) CreateAccount(email string) (*db.Account, error) {
	sa := new(storedAccount)
	insertAccount := "INSERT INTO accounts (email, created) VALUES ($1, $2) RETURNING *"
	err := c.Get(sa, insertAccount, email, time.Now())
	return sa.toAccount(), convertError(err)
}

func (c *conn) GetAccount(id int) (*db.Account, error) {
	sa := new(storedAccount)
	selectAccount := "SELECT * FROM accounts WHERE id=$1 LIMIT 1"
	err := c.Get(sa, selectAccount, id)
	return sa.toAccount(), convertError(err)
}

func (c *conn) GetAccountByEmail(email string) (*db.Account, error) {
	sa := new(storedAccount)
	selectAccount := "SELECT * FROM accounts WHERE email=$1 LIMIT 1"
	err := c.Get(sa, selectAccount, email)
	return sa.toAccount(), convertError(err)
}
