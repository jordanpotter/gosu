package postgres

import (
	"reflect"
	"testing"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func createTestChannel(t *testing.T, room db.Room) db.Channel {
	name := "device-name"
	channel, err := dbConn.CreateChannel(room.ID, name)
	if err != nil {
		t.Fatalf("Unexpected error during channel creation: %v", err)
	} else if channel.Name != name {
		t.Errorf("Mismatched channel name: %s != %s", channel.Name, name)
	} else if channel.Created.IsZero() {
		t.Errorf("Invalid timestamp: %v", channel.Created)
	}
	return channel
}

func TestChannelCreation(t *testing.T) {
	account := createTestAccount(t)
	room := createTestRoom(t, account)
	channel := createTestChannel(t, room)

	allChannels, err := dbConn.GetChannelsByRoom(room.ID)
	if err != nil {
		t.Fatalf("Unexpected error during channels retrieval: %v", err)
	} else if len(allChannels) != 1 {
		t.Errorf("Too many channels for room: %d", len(allChannels))
	} else if !reflect.DeepEqual(channel, allChannels[0]) {
		t.Errorf("Channels are not equal: %v != %v", channel, allChannels[0])
	}
}

func TestChannelDeletion(t *testing.T) {
	account := createTestAccount(t)
	room := createTestRoom(t, account)
	channel := createTestChannel(t, room)

	err := dbConn.DeleteChannelForRoom(channel.ID, room.ID)
	if err != nil {
		t.Errorf("Unexpected error during channel deletion: %v", err)
	}

	// Should allow repeated deletion
	err = dbConn.DeleteChannelForRoom(channel.ID, room.ID)
	if err != nil {
		t.Errorf("Unexpected error during channel deletion: %v", err)
	}

	allChannels, err := dbConn.GetChannelsByRoom(room.ID)
	if err != nil {
		t.Fatalf("Unexpected error during channels retrieval: %v", err)
	} else if len(allChannels) != 0 {
		t.Errorf("Too many channels for room: %d", len(allChannels))
	}
}
