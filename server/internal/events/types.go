package events

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"
)

const (
	accountDeviceCreatedType = "account.device.created"
	accountDeviceDeletedType = "account.device.deleted"

	roomChannelCreatedType      = "room.channel.created"
	roomChannelDeletedType      = "room.channel.deleted"
	roomMemberCreatedType       = "room.member.created"
	roomMemberDeletedType       = "room.member.deleted"
	roomMemberAdminUpdatedType  = "room.member.admin.updated"
	roomMemberBannedUpdatedType = "room.member.banned.updated"
)

type Type string

type Event interface {
	GetType() Type
}

func UnmarshalJSON(t Type, b []byte) (Event, error) {
	event, err := getEmptyEventForType(t)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, event)
	return event, err
}

func UnmarshalMsgpack(t Type, b []byte) (Event, error) {
	event, err := getEmptyEventForType(t)
	if err != nil {
		return nil, err
	}

	err = msgpack.Unmarshal(b, event)
	return event, err
}

func getEmptyEventForType(t Type) (Event, error) {
	switch t {
	case accountDeviceCreatedType:
		return new(AccountDeviceCreated), nil
	case accountDeviceDeletedType:
		return new(AccountDeviceDeleted), nil
	case roomChannelCreatedType:
		return new(RoomChannelCreated), nil
	case roomChannelDeletedType:
		return new(RoomChannelDeleted), nil
	case roomMemberCreatedType:
		return new(RoomMemberCreated), nil
	case roomMemberDeletedType:
		return new(RoomMemberDeleted), nil
	case roomMemberAdminUpdatedType:
		return new(RoomMemberAdminUpdated), nil
	case roomMemberBannedUpdatedType:
		return new(RoomMemberBannedUpdated), nil
	default:
		return nil, fmt.Errorf("unexpected type %s", t)
	}
}

type AccountDeviceCreated struct {
	AccountID  int       `json:"accountId" msgpack:"accountId"`
	DeviceID   int       `json:"deviceId" msgpack:"deviceId"`
	DeviceName string    `json:"deviceName" msgpack:"deviceName"`
	Created    time.Time `json:"created" msgpack:"created"`
}

func (adc *AccountDeviceCreated) GetType() Type {
	return accountDeviceCreatedType
}

type AccountDeviceDeleted struct {
	AccountID int `json:"accountId" msgpack:"accountId"`
	DeviceID  int `json"deviceId" msgpack:"deviceId"`
}

func (add *AccountDeviceDeleted) GetType() Type {
	return accountDeviceDeletedType
}

type RoomChannelCreated struct {
	RoomID      int       `json:"roomId" msgpack:"roomId"`
	ChannelID   int       `json:"channelId" msgpack:"channelId"`
	ChannelName string    `json:"roomName" msgpack:"roomName"`
	Created     time.Time `json:"created" msgpack:"created"`
}

func (rcc *RoomChannelCreated) GetType() Type {
	return roomChannelCreatedType
}

type RoomChannelDeleted struct {
	RoomID    int `json:"roomId" msgpack:"roomId"`
	ChannelID int `json:"channelId" msgpack:"channelId"`
}

func (rcd *RoomChannelDeleted) GetType() Type {
	return roomChannelDeletedType
}

type RoomMemberCreated struct {
	RoomID     int       `json:"roomId" msgpack:"roomId"`
	MemberID   int       `json:"memberId" msgpack:"memberId"`
	MemberName string    `json:"memberName", msgpack:"memberName"`
	Admin      bool      `json:"admin" msgpack:"admin"`
	Banned     bool      `json:"banned" msgpack:"banned"`
	Created    time.Time `json:"created" msgpack:"created"`
}

func (rmc *RoomMemberCreated) GetType() Type {
	return roomMemberCreatedType
}

type RoomMemberDeleted struct {
	RoomID   int `json:"roomId" msgpack:"roomId"`
	MemberID int `json:"memberId" msgpack:"memberId"`
}

func (rmd *RoomMemberDeleted) GetType() Type {
	return roomMemberDeletedType
}

type RoomMemberAdminUpdated struct {
	RoomID   int  `json:"roomId" msgpack:"roomId"`
	MemberID int  `json:"memberId" msgpack:"memberId"`
	Admin    bool `json:"admin" msgpack:"admin"`
}

func (rmau *RoomMemberAdminUpdated) GetType() Type {
	return roomMemberAdminUpdatedType
}

type RoomMemberBannedUpdated struct {
	RoomID   int  `json:"roomId" msgpack:"roomId"`
	MemberID int  `json:"memberId" msgpack:"memberId"`
	Banned   bool `json:"banned" msgpack:"banned"`
}

func (rmbu *RoomMemberBannedUpdated) GetType() Type {
	return roomMemberBannedUpdatedType
}
