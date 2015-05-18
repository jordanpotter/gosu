package rooms

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) Get(id string) (*db.Room, error) {
	sr := new(storedRoom)
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err := col.FindId(bson.ObjectIdHex(id)).One(sr)
	return handleStoredRoom(sr, err)
}

func (c *conn) GetByName(name string) (*db.Room, error) {
	sr := new(storedRoom)
	query := bson.M{"name": name}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err := col.Find(query).One(sr)
	return handleStoredRoom(sr, err)
}

func handleStoredRoom(sr *storedRoom, err error) (*db.Room, error) {
	if err == mgo.ErrNotFound {
		return nil, db.NotFoundError
	} else if err != nil {
		return nil, err
	}
	return sr.toRoom(), nil
}

func (c *conn) GetChannel(id, channelID string) (*db.Channel, error) {
	sr := new(storedRoom)
	query := bson.M{"_id": bson.ObjectIdHex(id), "channels.id": bson.ObjectIdHex(channelID)}
	projection := bson.M{"channels.$": 1}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err := col.Find(query).Select(projection).One(sr)
	return handleStoredChannelProjection(sr, err)
}

func handleStoredChannelProjection(sr *storedRoom, err error) (*db.Channel, error) {
	if err == mgo.ErrNotFound {
		return nil, db.NotFoundError
	} else if err != nil {
		return nil, err
	} else if len(sr.Channels) < 1 {
		return nil, db.NotFoundError
	}
	return sr.Channels[0].toChannel(), nil
}

func (c *conn) GetMember(id, memberID string) (*db.Member, error) {
	sr := new(storedRoom)
	query := bson.M{"_id": bson.ObjectIdHex(id), "members.id": bson.ObjectIdHex(memberID)}
	projection := bson.M{"members.$": 1}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err := col.Find(query).Select(projection).One(sr)
	return handleStoredMemberProjection(sr, err)
}

func (c *conn) GetMemberByAccount(id, accountID string) (*db.Member, error) {
	sr := new(storedRoom)
	query := bson.M{"_id": bson.ObjectIdHex(id), "members.accountID": bson.ObjectIdHex(accountID)}
	projection := bson.M{"members.$": 1}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err := col.Find(query).Select(projection).One(sr)
	return handleStoredMemberProjection(sr, err)
}

func handleStoredMemberProjection(sr *storedRoom, err error) (*db.Member, error) {
	if err == mgo.ErrNotFound {
		return nil, db.NotFoundError
	} else if err != nil {
		return nil, err
	} else if len(sr.Members) < 1 {
		return nil, db.NotFoundError
	}
	return sr.Members[0].toMember(), nil
}
