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

	query := bson.M{"email": email}
	device := bson.M{
		"id":           bson.NewObjectId(),
		"name":         deviceName,
		"passwordHash": dpHash,
		"created":      time.Now(),
	}
	onInsert := bson.M{"created": time.Now()}
	data := bson.M{"$push": bson.M{"devices": device}, "$setOnInsert": onInsert}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	_, err = col.Upsert(query, data)
	return err
}
