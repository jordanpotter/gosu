package sanitization

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/events"
)

const (
	RoomChannelCreatedName      = "roomChannelCreated"
	RoomChannelDeletedName      = "roomChannelDeleted"
	RoomMemberCreatedName       = "roomMemberCreated"
	RoomMemberDeletedName       = "roomMemberDeleted"
	RoomMemberAdminUpdatedName  = "roomMemberAdminUpdated"
	RoomMemberBannedUpdatedName = "roomMemberBannedUpdated"
)

type RoomChannelCreated struct {
	EventName   string    `json:"eventName"`
	RoomID      int       `json:"roomId"`
	ChannelID   int       `json:"channelId"`
	ChannelName string    `json:"channelName"`
	Created     time.Time `json:"created"`
	Timestamp   time.Time `json:"timestamp"`
}

func ToRoomChannelCreated(rcc events.RoomChannelCreated, timestamp time.Time) RoomChannelCreated {
	return RoomChannelCreated{
		EventName:   RoomChannelCreatedName,
		RoomID:      rcc.RoomID,
		ChannelID:   rcc.ChannelID,
		ChannelName: rcc.ChannelName,
		Created:     rcc.Created,
		Timestamp:   timestamp,
	}
}

type RoomChannelDeleted struct {
	EventName string    `json:"eventName"`
	RoomID    int       `json:"roomId"`
	ChannelID int       `json:"channelId"`
	Timestamp time.Time `json:"timestamp"`
}

func ToRoomChannelDeleted(rcd events.RoomChannelDeleted, timestamp time.Time) RoomChannelDeleted {
	return RoomChannelDeleted{
		EventName: RoomChannelDeletedName,
		RoomID:    rcd.RoomID,
		ChannelID: rcd.ChannelID,
		Timestamp: timestamp,
	}
}

type RoomMemberCreated struct {
	EventName  string    `json:"eventName"`
	RoomID     int       `json:"roomId"`
	MemberID   int       `json:"memberId"`
	MemberName string    `json:"memberName"`
	Admin      bool      `json:"admin"`
	Banned     bool      `json:"banned"`
	Created    time.Time `json:"created"`
	Timestamp  time.Time `json:"timestamp"`
}

func ToRoomMemberCreated(rmc events.RoomMemberCreated, timestamp time.Time) RoomMemberCreated {
	return RoomMemberCreated{
		EventName:  RoomMemberCreatedName,
		RoomID:     rmc.RoomID,
		MemberID:   rmc.MemberID,
		MemberName: rmc.MemberName,
		Admin:      rmc.Admin,
		Banned:     rmc.Banned,
		Created:    rmc.Created,
		Timestamp:  timestamp,
	}
}

type RoomMemberDeleted struct {
	EventName string    `json:"eventName"`
	RoomID    int       `json:"roomId"`
	MemberID  int       `json:"memberId"`
	Timestamp time.Time `json:"timestamp"`
}

func ToRoomMemberDeleted(rmd events.RoomMemberDeleted, timestamp time.Time) RoomMemberDeleted {
	return RoomMemberDeleted{
		EventName: RoomMemberDeletedName,
		RoomID:    rmd.RoomID,
		MemberID:  rmd.MemberID,
		Timestamp: timestamp,
	}
}

type RoomMemberAdminUpdated struct {
	EventName string    `json:"eventName"`
	RoomID    int       `json:"roomId"`
	MemberID  int       `json:"memberId"`
	Admin     bool      `json:"admin"`
	Timestamp time.Time `json:"timestamp"`
}

func ToRoomMemberAdminUpdated(rmau events.RoomMemberAdminUpdated, timestamp time.Time) RoomMemberAdminUpdated {
	return RoomMemberAdminUpdated{
		EventName: RoomMemberAdminUpdatedName,
		RoomID:    rmau.RoomID,
		MemberID:  rmau.MemberID,
		Admin:     rmau.Admin,
		Timestamp: timestamp,
	}
}

type RoomMemberBannedUpdated struct {
	EventName string    `json:"eventName"`
	RoomID    int       `json:"roomId"`
	MemberID  int       `json:"memberId"`
	Banned    bool      `json:"banned"`
	Timestamp time.Time `json:"timestamp"`
}

func ToRoomMemberBannedUpdated(rmau events.RoomMemberBannedUpdated, timestamp time.Time) RoomMemberBannedUpdated {
	return RoomMemberBannedUpdated{
		EventName: RoomMemberBannedUpdatedName,
		RoomID:    rmau.RoomID,
		MemberID:  rmau.MemberID,
		Banned:    rmau.Banned,
		Timestamp: timestamp,
	}
}
