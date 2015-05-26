package db

import (
	"time"
)

type RoomsConn interface {
	Create(name, password, adminAccountID, adminName string) (*Room, error)
	Get(id string) (*Room, error)
	GetByName(name string) (*Room, error)
	Delete(id string) error

	AddChannel(id, channelName string) (*Channel, error)
	GetChannel(id, channelID string) (*Channel, error)
	RemoveChannel(id, channelID string) error

	AddMember(id, accountID, memberName string) (*Member, error)
	GetMember(id, memberID string) (*Member, error)
	GetMemberByAccount(id, accountID string) (*Member, error)
	SetMemberAdmin(id, memberID string, admin bool) error
	SetMemberBanned(id, memberID string, banned bool) error
	RemoveMember(id, memberID string) error
}

type Room struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	PasswordHash []byte    `json:"-"` // For security, don't make public
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
	AccountID   string    `json:"-"` // For security, don't make public
	Name        string    `json:"name"`
	ChannelName string    `json:"channelName"`
	Admin       bool      `json:"admin"`
	Banned      bool      `json:"banned"`
	Created     time.Time `json:"created"`
}
