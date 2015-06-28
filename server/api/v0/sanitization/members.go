package sanitization

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type Member struct {
	ID        int       `json:"id"`
	RoomID    int       `json:"roomId"`
	ChannelID int       `json:"channelId"`
	Name      string    `json:"name"`
	Admin     bool      `json:"admin"`
	Banned    bool      `json:"banned"`
	Created   time.Time `json:"created"`
}

func ToMember(dbMember db.Member) Member {
	return Member{
		ID:        dbMember.ID,
		RoomID:    dbMember.RoomID,
		ChannelID: dbMember.ChannelID,
		Name:      dbMember.Name,
		Admin:     dbMember.Admin,
		Banned:    dbMember.Banned,
		Created:   dbMember.Created,
	}
}

func ToMembers(dbMembers []db.Member) []Member {
	members := make([]Member, 0, len(dbMembers))
	for _, dbMember := range dbMembers {
		members = append(members, ToMember(dbMember))
	}
	return members
}
