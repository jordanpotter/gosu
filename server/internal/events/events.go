package events

import "time"

type Message struct {
	Event interface{}
	Err   error
}

type RoomChannelCreated struct {
	RoomID      string    `msgpack:"roomID"`
	ChannelID   string    `msgpack:"channelID"`
	ChannelName string    `msgpack:"roomName"`
	Created     time.Time `msgpack:"created"`
}

type RoomChannelDeleted struct {
	RoomID    string `msgpack:"roomID"`
	ChannelID string `msgpack:"channelID"`
}

type RoomMemberCreated struct {
	RoomID     string    `msgpack:"roomID"`
	MemberID   string    `msgpack:"memberID"`
	MemberName string    `msgpack:"memberName"`
	Admin      bool      `msgpack:"admin"`
	Banned     bool      `msgpack:"banned"`
	Created    time.Time `msgpack:"created"`
}

type RoomMemberDeleted struct {
	RoomID   string `msgpack:"roomID"`
	MemberID string `msgpack:"memberID"`
}

type RoomMemberAdminUpdated struct {
	RoomID   string `msgpack:"roomID"`
	MemberID string `msgpack:"memberID"`
	Admin    bool   `msgpack:"admin"`
}

type RoomMemberBannedUpdated struct {
	RoomID   string `msgpack:"roomID"`
	MemberID string `msgpack:"memberID"`
	Banned   bool   `msgpack:"banned"`
}
