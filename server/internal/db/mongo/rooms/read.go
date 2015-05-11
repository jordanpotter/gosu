package rooms

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) Get(name string) (*db.Room, error) {
	var sr storedRoom
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err := col.Find(bson.M{"name": name}).One(&sr)
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
