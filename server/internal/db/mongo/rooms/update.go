package rooms

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) AddChannel(id, channelName string) error {
	exists, err := c.doesChannelExist(id, channelName)
	if err != nil {
		return err
	} else if exists {
		return db.DuplicateError
	}

	channel := bson.M{
		"id":      bson.NewObjectId(),
		"name":    channelName,
		"created": time.Now(),
	}
	data := bson.M{"$push": bson.M{"channels": channel}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.UpdateId(bson.ObjectIdHex(id), data)
}

func (c *conn) RemoveChannel(id, channelID string) error {
	channel := bson.M{"id": bson.ObjectIdHex(channelID)}
	data := bson.M{"$pull": bson.M{"channels": channel}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.UpdateId(bson.ObjectIdHex(id), data)
}

func (c *conn) doesChannelExist(id, channelName string) (bool, error) {
	query := bson.M{"_id": bson.ObjectIdHex(id), "channels.name": channelName}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	num, err := col.Find(query).Count()
	return num > 0, err
}

func (c *conn) AddMember(id, accountID, memberName string) error {
	exists, err := c.doesMemberExist(id, accountID, memberName)
	if err != nil {
		return err
	} else if exists {
		return db.DuplicateError
	}

	member := bson.M{
		"id":        bson.NewObjectId(),
		"accountID": bson.ObjectIdHex(accountID),
		"name":      memberName,
		"created":   time.Now(),
	}
	data := bson.M{"$push": bson.M{"members": member}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.UpdateId(bson.ObjectIdHex(id), data)
}

func (c *conn) SetMemberAdmin(id, memberID string, admin bool) error {
	query := bson.M{"_id": bson.ObjectIdHex(id), "members.id": bson.ObjectIdHex(memberID)}
	data := bson.M{"$set": bson.M{"members.$.admin": admin}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(query, data)
}

func (c *conn) SetMemberBanned(id, memberID string, banned bool) error {
	query := bson.M{"_id": bson.ObjectIdHex(id), "members.id": bson.ObjectIdHex(memberID)}
	data := bson.M{"$set": bson.M{"members.$.banned": banned}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(query, data)
}

func (c *conn) RemoveMember(id, memberID string) error {
	member := bson.M{"id": bson.ObjectIdHex(memberID)}
	data := bson.M{"$pull": bson.M{"members": member}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.UpdateId(bson.ObjectIdHex(id), data)
}

func (c *conn) doesMemberExist(id, accountID, memberName string) (bool, error) {
	accountOrName := []bson.M{
		bson.M{"members.accountID": bson.ObjectIdHex(accountID)},
		bson.M{"members.name": memberName},
	}
	query := bson.M{"_id": bson.ObjectIdHex(id), "$or": accountOrName}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	num, err := col.Find(query).Count()
	return num > 0, err
}
