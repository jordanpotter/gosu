package accounts

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (c *conn) AddMembership(id, roomId, peerName string) error {
	membershipBson := bson.M{"roomId": bson.ObjectIdHex(roomId), "peerName": peerName, "created": time.Now()}
	dataBson := bson.M{"$push": bson.M{"memberships": membershipBson}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	_, err := col.UpsertId(bson.ObjectIdHex(id), dataBson)
	return err
}
