package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JordanPotter/gosu-server/internal/auth/password"
	"github.com/JordanPotter/gosu-server/internal/db"
)

const defaultChannelName = "Lobby"

type storedRoom struct {
	Id           bson.ObjectId   `bson:"_id,omitempty"`
	Name         string          `bson:"name"`
	PasswordHash []byte          `bson:"passwordHash"`
	Channels     []storedChannel `bson:"channels"`
	Created      time.Time       `bson:"created"`
}

type storedChannel struct {
	Id      uint8        `bson:"id"`
	Name    string       `bson:"name"`
	Peers   []storedPeer `bson:"peers"`
	Created time.Time    `bson:"created"`
}

type storedPeer struct {
	Id        uint8         `bson:"id"`
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
		Id:      sc.Id,
		Name:    sc.Name,
		Peers:   peers,
		Created: sc.Created,
	}
}

func (sp *storedPeer) toPeer() *db.Peer {
	return &db.Peer{
		Id:        sp.Id,
		AccountId: sp.AccountId.Hex(),
		Name:      sp.Name,
		Created:   sp.Created,
	}
}

func (c *conn) ensureRoomIndices() error {
	nameIndex := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     false,
	}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.EnsureIndex(nameIndex)
}

func (c *conn) CreateRoom(name, pwd string) error {
	pHash, err := password.ComputeHash(pwd)
	if err != nil {
		return err
	}

	sChannel := storedChannel{
		Id:      0,
		Name:    defaultChannelName,
		Created: time.Now(),
	}

	sRoom := storedRoom{
		Name:         name,
		PasswordHash: pHash,
		Channels:     []storedChannel{sChannel},
		Created:      time.Now(),
	}

	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err = col.Insert(&sRoom)
	if mgo.IsDup(err) {
		return db.DuplicateError
	}
	return err
}
