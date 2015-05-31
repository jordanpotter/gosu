package types

import "time"

const (
	roomChannelCreatedType      = "room.channel.created"
	roomChannelDeletedType      = "room.channel.deleted"
	roomMemberCreatedType       = "room.member.created"
	roomMemberDeletedType       = "room.member.deleted"
	roomMemberAdminUpdatedType  = "room.member.admin.updated"
	roomMemberBannedUpdatedType = "room.member.banned.updated"
)

type RoomChannelCreated struct {
	RoomID      string    `json:"roomID" msgpack:"roomID"`
	ChannelID   string    `json:"channelID" msgpack:"channelID"`
	ChannelName string    `json:"roomName" msgpack:"roomName"`
	Created     time.Time `json:"created" msgpack:"created"`
}

func (rcc *RoomChannelCreated) GetType() Type {
	return roomChannelCreatedType
}

type RoomChannelDeleted struct {
	RoomID    string `json:"roomID" msgpack:"roomID"`
	ChannelID string `json"channelID" msgpack:"channelID"`
}

func (rcd *RoomChannelDeleted) GetType() Type {
	return roomChannelDeletedType
}

type RoomMemberCreated struct {
	RoomID     string    `json:"roomID" msgpack:"roomID"`
	MemberID   string    `json:"memberID" msgpack:"memberID"`
	MemberName string    `json:"memberName", msgpack:"memberName"`
	Admin      bool      `json:"admin" msgpack:"admin"`
	Banned     bool      `json:"banned" msgpack:"banned"`
	Created    time.Time `json:"created" msgpack:"created"`
}

func (rmc *RoomMemberCreated) GetType() Type {
	return roomMemberCreatedType
}

type RoomMemberDeleted struct {
	RoomID   string `json:"roomID" msgpack:"roomID"`
	MemberID string `json:"memberID" msgpack:"memberID"`
}

func (rmd *RoomMemberDeleted) GetType() Type {
	return roomMemberDeletedType
}

type RoomMemberAdminUpdated struct {
	RoomID   string `json:"roomID" msgpack:"roomID"`
	MemberID string `json:"memberID" msgpack:"memberID"`
	Admin    bool   `json:"admin" msgpack:"admin"`
}

func (rmau *RoomMemberAdminUpdated) GetType() Type {
	return roomMemberAdminUpdatedType
}

type RoomMemberBannedUpdated struct {
	RoomID   string `json:"roomID" msgpack:"roomID"`
	MemberID string `json:"memberID" msgpack:"memberID"`
	Banned   bool   `json:"banned" msgpack:"banned"`
}

func (rmbu *RoomMemberBannedUpdated) GetType() Type {
	return roomMemberBannedUpdatedType
}
