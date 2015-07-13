package postgres

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func createTestRoom(t *testing.T, account db.Account) db.Room {
	name := fmt.Sprintf("room-%d", rand.Uint32())
	passwordHash := []byte("password-hash")
	adminName := "Admin"

	room, err := dbConn.CreateRoom(name, passwordHash, account.ID, adminName)
	if err != nil {
		t.Fatalf("Unexpected error during room creation: %v", err)
	} else if room.Name != name {
		t.Errorf("Mismatched name: %s != %s", room.Name, name)
	} else if !bytes.Equal(room.PasswordHash, passwordHash) {
		t.Errorf("Mismatched password hash: %v != %v", room.PasswordHash, passwordHash)
	} else if room.Created.IsZero() {
		t.Errorf("Invalid timestamp: %v", room.Created)
	}
	return room
}

func TestRoomCreation(t *testing.T) {
	account := createTestAccount(t)
	room1 := createTestRoom(t, account)
	room2, err := dbConn.GetRoom(room1.ID)
	if err != nil {
		t.Fatalf("Unexpected error during room retrieval by id: %v", err)
	} else if !reflect.DeepEqual(room1, room2) {
		t.Errorf("Rooms are not equaul, %v != %v", room1, room2)
	}
}
