package db

import (
	"time"
)

type RoomsConn interface {
	Create(name, password, adminName, adminAccountID string) error
	Get(name string) (*Room, error)
	Delete(name string) error

	AddChannel(name, channelName string) error
	RemoveChannel(name, channelName string) error

	AddMember(name, accountID, memberName string) error
	GetMember(name, accountID string) (*Member, error)
	SetMemberAdmin(name, accountID string, admin bool) error
	SetMemberBanned(name, accountID string, banned bool) error
	RemoveMember(name, accountID string) error
}

type Room struct {
	Name         string    `json:"name"`
	PasswordHash []byte    `json:"-"`
	Channels     []Channel `json:"channels"`
	Members      []Member  `json:"members"`
	Created      time.Time `json:"created"`
}

type Channel struct {
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

type Member struct {
	AccountID   string    `json:"accountID"`
	Name        string    `json:"name"`
	ChannelName string    `json:"channelName"`
	Admin       bool      `json:"admin"`
	Banned      bool      `json:"banned"`
	Created     time.Time `json:"created"`
}
