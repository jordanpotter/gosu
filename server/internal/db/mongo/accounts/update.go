package accounts

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (c *conn) AddMembership(id, roomId, peerName string) error {
	err := c.RemoveMembership(id, roomId)
	if err != nil {
		return err
	}

	membershipBson := bson.M{"roomId": bson.ObjectIdHex(roomId), "peerName": peerName, "created": time.Now()}
	dataBson := bson.M{"$push": bson.M{"memberships": membershipBson}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	return col.UpdateId(bson.ObjectIdHex(id), dataBson)
}

func (c *conn) SetMembershipAdmin(id, roomId string, admin bool) error {
	queryBson := bson.M{"id": bson.ObjectIdHex(id), "memberships.roomId": bson.ObjectIdHex(roomId)}
	dataBson := bson.M{"$set": bson.M{"memberships.$.admin": admin}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	return col.Update(queryBson, dataBson)
}

func (c *conn) SetMembershipBanned(id, roomId string, banned bool) error {
	queryBson := bson.M{"id": bson.ObjectIdHex(id), "memberships.roomId": bson.ObjectIdHex(roomId)}
	dataBson := bson.M{"$set": bson.M{"memberships.$.banned": banned}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	return col.Update(queryBson, dataBson)
}

func (c *conn) RemoveMembership(id, roomId string) error {
	matchBson := bson.M{"roomId": bson.ObjectIdHex(roomId)}
	dataBson := bson.M{"$pull": bson.M{"memberships": matchBson}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	return col.UpdateId(bson.ObjectIdHex(id), dataBson)
}
