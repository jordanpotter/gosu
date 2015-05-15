package rooms

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) Get(id string) (*db.Room, error) {
	var sr storedRoom
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err := col.FindId(bson.ObjectIdHex(id)).One(&sr)
	return handleStoredRoom(&sr, err)
}

func (c *conn) GetByName(name string) (*db.Room, error) {
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
