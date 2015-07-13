package postgres

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestRoomCreation(t *testing.T) {
	email := fmt.Sprintf("test-%d@email.com", rand.Uint32())
	account, err := dbConn.CreateAccount(email)
	if err != nil {
		t.Errorf("Unexpected error during account creation: %v", err)
	}

	name := fmt.Sprintf("room-%d", rand.Uint32())
	passwordHash := []byte("password-hash")
	adminName := "Admin"

	room1, err := dbConn.CreateRoom(name, passwordHash, account.ID, adminName)
	if err != nil {
		t.Errorf("Unexpected error during room creation: %v", err)
	} else if room1.Name != name {
		t.Errorf("Mismatched name, %s != %s", room1.Name, name)
	} else if !bytes.Equal(room1.PasswordHash, passwordHash) {
		t.Errorf("Mismatched password hash, %v != %v", room1.PasswordHash, passwordHash)
	} else if room1.Created.IsZero() {
		t.Errorf("Invalid timestamp, %v", room1.Created)
	}

	room2, err := dbConn.GetRoom(room1.ID)
	if err != nil {
		t.Errorf("Unexpected error during room retrieval by id: %v", err)
	} else if !reflect.DeepEqual(room1, room2) {
		t.Errorf("Rooms are not equaul, %v != %v", room1, room2)
	}
}
