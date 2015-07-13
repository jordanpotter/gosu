package postgres

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func createTestAccount(t *testing.T) db.Account {
	email := fmt.Sprintf("test-%d@email.com", rand.Uint32())
	account, err := dbConn.CreateAccount(email)
	if err != nil {
		t.Fatalf("Unexpected error during account creation: %v", err)
	} else if account.Email != email {
		t.Errorf("Mismatched email: %s != %s", account.Email, email)
	} else if account.Created.IsZero() {
		t.Errorf("Invalid timestamp: %v", account.Created)
	}
	return account
}

func TestAccountCreation(t *testing.T) {
	account1 := createTestAccount(t)

	account2, err := dbConn.GetAccount(account1.ID)
	if err != nil {
		t.Fatalf("Unexpected error during account retrieval by id: %v", err)
	} else if !reflect.DeepEqual(account1, account2) {
		t.Errorf("Accounts are not equaul: %v != %v", account1, account2)
	}

	account3, err := dbConn.GetAccountByEmail(account1.Email)
	if err != nil {
		t.Fatalf("Unexpected error during account retrieval by email: %v", err)
	} else if !reflect.DeepEqual(account1, account3) {
		t.Errorf("Accounts are not equal: %v != %v", account1, account3)
	}
}
