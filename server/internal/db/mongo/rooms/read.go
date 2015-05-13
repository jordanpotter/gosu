package rooms

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) Get(name string) (*db.Room, error) {
	var sr storedRoom
	query := bson.M{"name": name}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err := col.Find(query).One(&sr)
	return handleStoredRoom(&sr, err)
}

func handleStoredRoom(sr *storedRoom, err error) (*db.Room, error) {
	if err == mgo.ErrNotFound {
		return nil, db.NotFoundError
	} else if err != nil {
		return nil, err
	}
	return sr.toRoom(), nil
}

func (c *conn) GetMember(name, accountID string) (*db.Member, error) {
	room, err := c.Get(name)
	if err != nil {
		return nil, err
	}

	for _, member := range room.Members {
		if member.AccountID == accountID {
			return &member, nil
		}
	}
	return nil, db.NotFoundError
}
