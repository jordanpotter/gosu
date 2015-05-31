package types

import (
	"encoding/json"
	"fmt"

	"gopkg.in/vmihailenco/msgpack.v2"
)

func init() {
	fmt.Println("MOVE ALL OF THIS TO THE PARENT DIRECTORY (EVENTS) IF THERE ARE NO CIRCULAR DEPENDENCIES")
}

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
