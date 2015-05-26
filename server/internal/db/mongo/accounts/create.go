package accounts

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) Create(email, deviceName, devicePassword string) (*db.Account, error) {
	dpHash, err := password.ComputeHash(devicePassword)
	if err != nil {
		return nil, err
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
	change := mgo.Change{
		Update:    data,
		Upsert:    true,
		ReturnNew: true,
	}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	sa := new(storedAccount)
	_, err = col.Find(query).Apply(change, sa)
	return sa.toAccount(), err
}
