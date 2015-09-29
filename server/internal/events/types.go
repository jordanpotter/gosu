package events

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jordanpotter/gosu/Godeps/_workspace/src/gopkg.in/vmihailenco/msgpack.v2"
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
	switch t {
	case accountDeviceCreatedType:
		var adc AccountDeviceCreated
		err := json.Unmarshal(b, &adc)
		return adc, err
	case accountDeviceDeletedType:
		var add AccountDeviceDeleted
		err := json.Unmarshal(b, &add)
		return add, err
	case roomChannelCreatedType:
		var rcc RoomChannelCreated
		err := json.Unmarshal(b, &rcc)
		return rcc, err
	case roomChannelDeletedType:
		var rcd RoomChannelDeleted
		err := json.Unmarshal(b, &rcd)
		return rcd, err
	case roomMemberCreatedType:
		var rmc RoomMemberCreated
		err := json.Unmarshal(b, &rmc)
		return rmc, err
	case roomMemberDeletedType:
		var rmd RoomMemberDeleted
		err := json.Unmarshal(b, &rmd)
		return rmd, err
	case roomMemberAdminUpdatedType:
		var rmau RoomMemberAdminUpdated
		err := json.Unmarshal(b, &rmau)
		return rmau, err
	case roomMemberBannedUpdatedType:
		var rmbu RoomMemberBannedUpdated
		err := json.Unmarshal(b, &rmbu)
		return rmbu, err
	default:
		return nil, fmt.Errorf("unexpected type %s", t)
	}
}

func UnmarshalMsgpack(t Type, b []byte) (Event, error) {
	switch t {
	case accountDeviceCreatedType:
		var adc AccountDeviceCreated
		err := msgpack.Unmarshal(b, &adc)
		return adc, err
	case accountDeviceDeletedType:
		var add AccountDeviceDeleted
		err := msgpack.Unmarshal(b, &add)
		return add, err
	case roomChannelCreatedType:
		var rcc RoomChannelCreated
		err := msgpack.Unmarshal(b, &rcc)
		return rcc, err
	case roomChannelDeletedType:
		var rcd RoomChannelDeleted
		err := msgpack.Unmarshal(b, &rcd)
		return rcd, err
	case roomMemberCreatedType:
		var rmc RoomMemberCreated
		err := msgpack.Unmarshal(b, &rmc)
		return rmc, err
	case roomMemberDeletedType:
		var rmd RoomMemberDeleted
		err := msgpack.Unmarshal(b, &rmd)
		return rmd, err
	case roomMemberAdminUpdatedType:
		var rmau RoomMemberAdminUpdated
		err := msgpack.Unmarshal(b, &rmau)
		return rmau, err
	case roomMemberBannedUpdatedType:
		var rmbu RoomMemberBannedUpdated
		err := msgpack.Unmarshal(b, &rmbu)
		return rmbu, err
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

func (adc AccountDeviceCreated) GetType() Type {
	return accountDeviceCreatedType
}

type AccountDeviceDeleted struct {
	AccountID int `json:"accountId" msgpack:"accountId"`
	DeviceID  int `json:"deviceId" msgpack:"deviceId"`
}

func (add AccountDeviceDeleted) GetType() Type {
	return accountDeviceDeletedType
}

type RoomChannelCreated struct {
	RoomID      int       `json:"roomId" msgpack:"roomId"`
	ChannelID   int       `json:"channelId" msgpack:"channelId"`
	ChannelName string    `json:"roomName" msgpack:"roomName"`
	Created     time.Time `json:"created" msgpack:"created"`
}

func (rcc RoomChannelCreated) GetType() Type {
	return roomChannelCreatedType
}

type RoomChannelDeleted struct {
	RoomID    int `json:"roomId" msgpack:"roomId"`
	ChannelID int `json:"channelId" msgpack:"channelId"`
}

func (rcd RoomChannelDeleted) GetType() Type {
	return roomChannelDeletedType
}

type RoomMemberCreated struct {
	RoomID     int       `json:"roomId" msgpack:"roomId"`
	MemberID   int       `json:"memberId" msgpack:"memberId"`
	MemberName string    `json:"memberName" msgpack:"memberName"`
	Admin      bool      `json:"admin" msgpack:"admin"`
	Banned     bool      `json:"banned" msgpack:"banned"`
	Created    time.Time `json:"created" msgpack:"created"`
}

func (rmc RoomMemberCreated) GetType() Type {
	return roomMemberCreatedType
}

type RoomMemberDeleted struct {
	RoomID   int `json:"roomId" msgpack:"roomId"`
	MemberID int `json:"memberId" msgpack:"memberId"`
}

func (rmd RoomMemberDeleted) GetType() Type {
	return roomMemberDeletedType
}

type RoomMemberAdminUpdated struct {
	RoomID   int  `json:"roomId" msgpack:"roomId"`
	MemberID int  `json:"memberId" msgpack:"memberId"`
	Admin    bool `json:"admin" msgpack:"admin"`
}

func (rmau RoomMemberAdminUpdated) GetType() Type {
	return roomMemberAdminUpdatedType
}

type RoomMemberBannedUpdated struct {
	RoomID   int  `json:"roomId" msgpack:"roomId"`
	MemberID int  `json:"memberId" msgpack:"memberId"`
	Banned   bool `json:"banned" msgpack:"banned"`
}

func (rmbu RoomMemberBannedUpdated) GetType() Type {
	return roomMemberBannedUpdatedType
}
