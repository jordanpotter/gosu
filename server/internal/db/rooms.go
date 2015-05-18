package db

import (
	"time"
)

type RoomsConn interface {
	Create(name, password, adminAccountID, adminName string) error
	Get(id string) (*Room, error)
	GetByName(name string) (*Room, error)
	Delete(id string) error

	AddChannel(id, channelName string) error
	GetChannel(id, channelID string) (*Channel, error)
	RemoveChannel(id, channelID string) error

	AddMember(id, accountID, memberName string) error
	GetMember(id, memberID string) (*Member, error)
	GetMemberByAccount(id, accountID string) (*Member, error)
	SetMemberAdmin(id, memberID string, admin bool) error
	SetMemberBanned(id, memberID string, banned bool) error
	RemoveMember(id, memberID string) error
	RemoveMemberByAccount(id, accountID string) error
}

type Room struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	PasswordHash []byte    `json:"-"`
	Channels     []Channel `json:"channels"`
	Members      []Member  `json:"members"`
	Created      time.Time `json:"created"`
}

type Channel struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

type Member struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"accountID"`
	Name        string    `json:"name"`
	ChannelName string    `json:"channelName"`
	Admin       bool      `json:"admin"`
	Banned      bool      `json:"banned"`
	Created     time.Time `json:"created"`
}
