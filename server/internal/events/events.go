package events

const (
	RoomCreated = "room.created"
	RoomDeleted = "room.deleted"

	RoomChannelCreated = "room.channel.created"
	RoomChannelDeleted = "room.channel.deleted"

	RoomMemberCreated       = "room.member.created"
	RoomMemberDeleted       = "room.member.deleted"
	RoomMemberUpdatedAdmin  = "room.member.updated.admin"
	RoomMemberUpdatedBanned = "room.member.updated.banned"
)

type Message struct {
	Data []byte
	Err  error
}
