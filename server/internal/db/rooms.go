package db

import (
	"time"
)

type RoomsConn interface {
	Create(name, password string) error
	Get(id string) (*Room, error)
	GetByName(name string) (*Room, error)
	Delete(id string) error

	// CreateChannel(id, name string) (*Channel, error)
	// DeleteChannel(id channelID string) error

	// AddMember(id, name string) error
	// SetMemberAdmin(id, mDmberId string, admin bool) error
	// SetMemberBanned(id, memberId string, banned bool) error
	// RemoveMember(id, memberID string) error
}

type Room struct {
	ID           string
	Name         string
	PasswordHash []byte
	Channels     []Channel
	Members      []Member
	Created      time.Time
}

type Channel struct {
	ID      string
	Name    string
	Created time.Time
}

type Member struct {
	ID        string
	AccountID string
	Name      string
	ChannelID string
	Admin     bool
	Banned    bool
	Created   time.Time
}
