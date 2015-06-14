package db

import "io"

type Conn interface {
	AccountsConn
	DevicesConn
	RoomsConn
	ChannelsConn
	MembersConn
	io.Closer
}

type AccountsConn interface {
	CreateAccount(email string) (*Account, error)
	GetAccount(id int) (*Account, error)
	GetAccountByEmail(email string) (*Account, error)
}

type DevicesConn interface {
	CreateDevice(accountID int, deviceName string, devicePasswordHash []byte) (*Device, error)
	GetDevicesByAccount(accountID int) ([]Device, error)
	DeleteDevice(id int) error
}

type RoomsConn interface {
	CreateRoom(name string, passwordHash []byte, adminAccountID int, adminName string) (*Room, error)
	GetRoom(id int) (*Room, error)
	GetRoomByName(name string) (*Room, error)
}

type ChannelsConn interface {
	CreateChannel(roomID int, name string) (*Channel, error)
	GetChannel(id int) (*Channel, error)
	GetChannelsByRoom(roomID int) ([]Channel, error)
	DeleteChannel(id int) error
}

type MembersConn interface {
	CreateMember(accountID, roomID int, name string) (*Member, error)
	GetMember(id int) (*Member, error)
	GetMembersByAccount(accountID int) ([]Member, error)
	GetMembersByRoom(roomID int) ([]Member, error)
	SetMemberAdmin(id int, admin bool) (*Member, error)
	SetMemberBanned(id int, banned bool) (*Member, error)
	DeleteMember(id int) error
}
