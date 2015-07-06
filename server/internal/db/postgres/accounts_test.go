package postgres

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestAccountCreation(t *testing.T) {
	email := fmt.Sprintf("test-%d@email.com", rand.Uint32())
	account1, err := dbConn.CreateAccount(email)
	if err != nil {
		t.Errorf("Unexpected error during account creation: %v", err)
	} else if account1.Email != email {
		t.Errorf("Mismatched email, %s != %s", account1.Email, email)
	}

	account2, err := dbConn.GetAccount(account1.ID)
	if err != nil {
		t.Errorf("Unexpected error during account retrieval by id: %v", err)
	} else if !reflect.DeepEqual(account1, account2) {
		t.Errorf("Accounts are not equaul, %v != %v", account1, account2)
	}

	account3, err := dbConn.GetAccountByEmail(email)
	if err != nil {
		t.Errorf("Unexpected error during account retrieval by email: %v", err)
	} else if !reflect.DeepEqual(account1, account3) {
		t.Errorf("Accounts are not equal, %v != %v", account1, account3)
	}
}
