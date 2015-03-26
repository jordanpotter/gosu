package db

import (
	"time"
)

type RoomsConn interface {
	Create(name, password string) error
	// GetRoom(id string) (*Room, error)
	// DeleteRoom(id string) error

	// ConnectToRoom(id, accountId string) (*Room, error)
	// DisconnectFromRoom(id, accountId string) error

	// CreateChannel(id, name string) (*Channel, error)
	// DeleteChannel(id string, channelId uint8) error

	// MoveToChannel(id string, channelId uint8, accountId string) error
}

type Room struct {
	Id           string
	Name         string
	PasswordHash []byte
	Channels     []Channel
	Created      time.Time
}

type Channel struct {
	Id      string
	Name    string
	Peers   []Peer
	Created time.Time
}

type Peer struct {
	Id        string
	AccountId string
	Name      string
	Created   time.Time
}
