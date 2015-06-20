package sanitization

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type Account struct {
	ID      int       `json:"id"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
}

func ToAccount(dbAccount *db.Account) *Account {
	return &Account{
		ID:      dbAccount.ID,
		Email:   dbAccount.Email,
		Created: dbAccount.Created,
	}
}

func ToAccounts(dbAccounts []db.Account) []Account {
	accounts := make([]Account, 0, len(dbAccounts))
	for _, dbAccount := range dbAccounts {
		accounts = append(accounts, *ToAccount(&dbAccount))
	}
	return accounts
}
