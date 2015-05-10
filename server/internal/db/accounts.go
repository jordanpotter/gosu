package db

import (
	"time"
)

type AccountsConn interface {
	Create(email, deviceName, devicePassword string) error
	Get(id string) (*Account, error)
	GetByEmail(email string) (*Account, error)
	GetByMembershipId(membershipId string) (*Account, error)

	AddMembership(id, roomId, peerName string) error
	SetMembershipAdmin(id, roomId string, admin bool) error
	SetMembershipBanned(id, roomId string, banned bool) error
	RemoveMembership(id, roomId string) error
}

type Account struct {
	Id          string
	Email       string
	Devices     []Device
	Memberships []Membership
	Created     time.Time
}

type Device struct {
	Name         string
	PasswordHash []byte
	Created      time.Time
}

type Membership struct {
	Id       string
	RoomId   string
	PeerName string
	Admin    bool
	Banned   bool
	Created  time.Time
}
