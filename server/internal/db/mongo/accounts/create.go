package accounts

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/auth/password"
)

func (c *conn) Create(email, deviceName, devicePassword string) error {
	dpHash, err := password.ComputeHash(devicePassword)
	if err != nil {
		return err
	}

	findBson := bson.M{"email": email}
	deviceBson := bson.M{"name": deviceName, "passwordHash": dpHash, "created": time.Now()}
	dataBson := bson.M{"$push": bson.M{"devices": deviceBson}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	_, err = col.Upsert(findBson, dataBson)
	return err
}
