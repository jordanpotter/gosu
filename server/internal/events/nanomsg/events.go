package nanomsg

import (
	"fmt"

	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/jordanpotter/gosu/server/internal/events"
)

const (
	roomChannelCreated      = "room.channel.created"
	roomChannelDeleted      = "room.channel.deleted"
	roomMemberCreated       = "room.member.created"
	roomMemberDeleted       = "room.member.deleted"
	roomMemberAdminUpdated  = "room.member.admin.updated"
	roomMemberBannedUpdated = "room.member.banned.updated"
)

type Message struct {
	Type       string `msgpack:"type"`
	EventBytes []byte `msgpack:eventBytes`
}

func newMessage(event interface{}) (*Message, error) {
	t, err := getEventType(event)
	if err != nil {
		return nil, err
	}

	b, err := msgpack.Marshal(event)
	return &Message{Type: t, EventBytes: b}, err
}

func (m *Message) getEvent() (interface{}, error) {
	var out interface{}
	switch m.Type {
	case roomChannelCreated:
		out = new(events.RoomChannelCreated)
	case roomChannelDeleted:
		out = new(events.RoomChannelDeleted)
	case roomMemberCreated:
		out = new(events.RoomMemberCreated)
	case roomMemberDeleted:
		out = new(events.RoomMemberDeleted)
	case roomMemberAdminUpdated:
		out = new(events.RoomMemberAdminUpdated)
	case roomMemberBannedUpdated:
		out = new(events.RoomMemberBannedUpdated)
	default:
		return nil, fmt.Errorf("unexpected type %s", m.Type)
	}

	err := msgpack.Unmarshal(m.EventBytes, out)
	return out, err
}

func getEventType(event interface{}) (string, error) {
	switch t := event.(type) {
	case events.RoomChannelCreated:
		return roomChannelCreated, nil
	case events.RoomChannelDeleted:
		return roomChannelDeleted, nil
	case events.RoomMemberCreated:
		return roomMemberCreated, nil
	case events.RoomMemberDeleted:
		return roomMemberDeleted, nil
	case events.RoomMemberAdminUpdated:
		return roomMemberAdminUpdated, nil
	case events.RoomMemberBannedUpdated:
		return roomMemberBannedUpdated, nil
	default:
		return "", fmt.Errorf("unexpected type %T", t)
	}
}
