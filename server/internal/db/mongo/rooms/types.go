package rooms

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedRoom struct {
	Name         string          `bson:"name"`
	PasswordHash []byte          `bson:"passwordHash"`
	Channels     []storedChannel `bson:"channels"`
	Members      []storedMember  `bson:"members"`
	Created      time.Time       `bson:"created"`
}

type storedChannel struct {
	Name    string    `bson:"name"`
	Created time.Time `bson:"created"`
}

type storedMember struct {
	Name        string        `bson:"name"`
	AccountID   bson.ObjectId `bson:"accountID"`
	ChannelName string        `bson:"channelName"`
	Admin       bool          `bson:"admin"`
	Banned      bool          `bson:"banned"`
	Created     time.Time     `bson:"created"`
}

func (sr *storedRoom) toRoom() *db.Room {
	channels := make([]db.Channel, len(sr.Channels))
	for _, sChannel := range sr.Channels {
		channels = append(channels, *sChannel.toChannel())
	}

	members := make([]db.Member, len(sr.Members))
	for _, sMember := range sr.Members {
		members = append(members, *sMember.toMember())
	}

	return &db.Room{
		Name:         sr.Name,
		PasswordHash: sr.PasswordHash,
		Channels:     channels,
		Members:      members,
		Created:      sr.Created,
	}
}

func (sc *storedChannel) toChannel() *db.Channel {
	return &db.Channel{
		Name:    sc.Name,
		Created: sc.Created,
	}
}

func (sm *storedMember) toMember() *db.Member {
	return &db.Member{
		Name:        sm.Name,
		AccountID:   sm.AccountID.Hex(),
		ChannelName: sm.ChannelName,
		Admin:       sm.Admin,
		Banned:      sm.Banned,
		Created:     sm.Created,
	}
}
