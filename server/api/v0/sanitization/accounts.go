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

func SanitizeAccount(dbAccount *db.Account) *Account {
	return &Account{
		ID:      dbAccount.ID,
		Email:   dbAccount.Email,
		Created: dbAccount.Created,
	}
}

func SanitizeAccounts(dbAccounts []db.Account) []Account {
	accounts := make([]Account, 0, len(dbAccounts))
	for _, dbAccount := range dbAccounts {
		accounts = append(accounts, *SanitizeAccount(&dbAccount))
	}
	return accounts
}
