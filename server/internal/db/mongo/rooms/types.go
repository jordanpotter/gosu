package rooms

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/JordanPotter/gosu-server/server/internal/db"
)

type storedRoom struct {
	Id           bson.ObjectId   `bson:"_id,omitempty"`
	Name         string          `bson:"name"`
	PasswordHash []byte          `bson:"passwordHash"`
	Channels     []storedChannel `bson:"channels"`
	Created      time.Time       `bson:"created"`
}

type storedChannel struct {
	Id      bson.ObjectId `bson:"id"`
	Name    string        `bson:"name"`
	Peers   []storedPeer  `bson:"peers"`
	Created time.Time     `bson:"created"`
}

type storedPeer struct {
	Id        bson.ObjectId `bson:"id"`
	AccountId bson.ObjectId `bson:"accountId"`
	Name      string        `bson:"name"`
	Created   time.Time     `bson:"created"`
}

func (sr *storedRoom) toRoom() *db.Room {
	channels := make([]db.Channel, len(sr.Channels))
	for _, sChannel := range sr.Channels {
		channels = append(channels, *sChannel.toChannel())
	}

	return &db.Room{
		Id:           sr.Id.Hex(),
		Name:         sr.Name,
		PasswordHash: sr.PasswordHash,
		Channels:     channels,
		Created:      sr.Created,
	}
}

func (sc *storedChannel) toChannel() *db.Channel {
	peers := make([]db.Peer, len(sc.Peers))
	for _, sPeer := range sc.Peers {
		peers = append(peers, *sPeer.toPeer())
	}

	return &db.Channel{
		Id:      sc.Id.Hex(),
		Name:    sc.Name,
		Peers:   peers,
		Created: sc.Created,
	}
}

func (sp *storedPeer) toPeer() *db.Peer {
	return &db.Peer{
		Id:        sp.Id.Hex(),
		AccountId: sp.AccountId.Hex(),
		Name:      sp.Name,
		Created:   sp.Created,
	}
}
