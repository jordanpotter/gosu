package rooms

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) AddChannel(name, channelName string) error {
	exists, err := c.doesChannelExist(name, channelName)
	if err != nil {
		return err
	} else if exists {
		return db.DuplicateError
	}

	query := bson.M{"name": name}
	channel := bson.M{"name": channelName, "created": time.Now()}
	data := bson.M{"$push": bson.M{"channels": channel}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(query, data)
}

func (c *conn) RemoveChannel(name, channelName string) error {
	exists, err := c.doesChannelExist(name, channelName)
	if err != nil {
		return err
	} else if !exists {
		return db.NotFoundError
	}

	query := bson.M{"name": name}
	channel := bson.M{"name": channelName}
	data := bson.M{"$pull": bson.M{"channels": channel}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(query, data)
}

func (c *conn) doesChannelExist(name, channelName string) (bool, error) {
	query := bson.M{"name": name, "channels.name": channelName}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	num, err := col.Find(query).Count()
	return num > 0, err
}

func (c *conn) AddMember(name, accountID, memberName string) error {
	exists, err := c.doesMemberExist(name, accountID, memberName)
	if err != nil {
		return err
	} else if exists {
		return db.DuplicateError
	}

	query := bson.M{"name": name}
	member := bson.M{"accountID": bson.ObjectIdHex(accountID), "name": memberName, "created": time.Now()}
	data := bson.M{"$push": bson.M{"members": member}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(query, data)
}

func (c *conn) SetMemberAdmin(name, accountID string, admin bool) error {
	query := bson.M{"name": name, "members.accountID": accountID}
	data := bson.M{"$set": bson.M{"members.$.admin": admin}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(query, data)
}

func (c *conn) SetMemberBanned(name, accountID string, banned bool) error {
	query := bson.M{"name": name, "members.accountID": accountID}
	data := bson.M{"$set": bson.M{"members.$.banned": banned}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(query, data)
}

func (c *conn) RemoveMember(name, accountID string) error {
	exists, err := c.doesMemberExist(name, accountID, "")
	if err != nil {
		return err
	} else if !exists {
		return db.NotFoundError
	}

	query := bson.M{"name": name}
	member := bson.M{"name": accountID}
	data := bson.M{"$pull": bson.M{"members": member}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(query, data)
}

func (c *conn) doesMemberExist(name, accountID, memberName string) (bool, error) {
	accountOrName := []bson.M{
		bson.M{"members.accountID": bson.ObjectIdHex(accountID)},
		bson.M{"members.name": memberName},
	}
	query := bson.M{"name": name, "$or": accountOrName}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	num, err := col.Find(query).Count()
	return num > 0, err
}
