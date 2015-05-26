package rooms

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/db"
)

const defaultChannelName = "Lobby"

func (c *conn) Create(name, passwd, adminAccountID, adminName string) (*db.Room, error) {
	pHash, err := password.ComputeHash(passwd)
	if err != nil {
		return nil, err
	}

	sChannel := storedChannel{
		ID:      bson.NewObjectId(),
		Name:    defaultChannelName,
		Created: time.Now(),
	}

	sMember := storedMember{
		ID:        bson.NewObjectId(),
		AccountID: bson.ObjectIdHex(adminAccountID),
		Name:      adminName,
		Admin:     true,
		Created:   time.Now(),
	}

	sRoom := storedRoom{
		ID:           bson.NewObjectId(),
		Name:         name,
		PasswordHash: pHash,
		Channels:     []storedChannel{sChannel},
		Members:      []storedMember{sMember},
		Created:      time.Now(),
	}

	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err = col.Insert(&sRoom)
	if mgo.IsDup(err) {
		return nil, db.DuplicateError
	}
	return sRoom.toRoom(), err
}
