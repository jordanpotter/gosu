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

	channelBson := bson.M{"name": channelName, "created": time.Now()}
	dataBson := bson.M{"$push": bson.M{"channels": channelBson}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(bson.M{"name": name}, dataBson)
}

func (c *conn) RemoveChannel(name, channelName string) error {
	channelBson := bson.M{"name": channelName}
	dataBson := bson.M{"$pull": bson.M{"channels": channelBson}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(bson.M{"name": name}, dataBson)
}

func (c *conn) doesChannelExist(name, channelName string) (bool, error) {
	queryBson := bson.M{"name": name, "channels.name": channelName}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	num, err := col.Find(queryBson).Count()
	return num > 0, err
}

func (c *conn) AddMember(name, memberName, accountID string) error {
	exists, err := c.doesMemberExist(name, memberName, accountID)
	if err != nil {
		return err
	} else if exists {
		return db.DuplicateError
	}

	memberBson := bson.M{"name": memberName, "accountId": bson.ObjectIdHex(accountID), "created": time.Now()}
	dataBson := bson.M{"$push": bson.M{"members": memberBson}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(bson.M{"name": name}, dataBson)
}

func (c *conn) SetMemberAdmin(name, memberName string, admin bool) error {
	queryBson := bson.M{"name": name, "members.name": memberName}
	dataBson := bson.M{"$set": bson.M{"members.$.admin": admin}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(queryBson, dataBson)
}

func (c *conn) SetMemberBanned(name, memberName string, banned bool) error {
	queryBson := bson.M{"name": name, "members.name": memberName}
	dataBson := bson.M{"$set": bson.M{"members.$.banned": banned}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(queryBson, dataBson)
}

func (c *conn) RemoveMember(name, memberName string) error {
	memberBson := bson.M{"name": memberName}
	dataBson := bson.M{"$pull": bson.M{"members": memberBson}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.Update(bson.M{"name": name}, dataBson)
}

func (c *conn) doesMemberExist(name, memberName, accountID string) (bool, error) {
	nameOrAccountBson := []bson.M{
		bson.M{"members.name": memberName},
		bson.M{"members.accountId": bson.ObjectIdHex(accountID)},
	}
	queryBson := bson.M{"name": name, "$or": nameOrAccountBson}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	num, err := col.Find(queryBson).Count()
	return num > 0, err
}
