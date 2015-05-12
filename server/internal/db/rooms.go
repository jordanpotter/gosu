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

	AddMember(name, memberName, accountID string) error
	SetMemberAdmin(name, memberName string, admin bool) error
	SetMemberBanned(name, memberName string, banned bool) error
	RemoveMember(name, memberName string) error
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
	Name        string    `json:"name"`
	AccountID   string    `json:"-"`
	ChannelName string    `json:"channelName"`
	Admin       bool      `json:"admin"`
	Banned      bool      `json:"banned"`
	Created     time.Time `json:"created"`
}
