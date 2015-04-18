package rooms

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/db"
)

const defaultChannelName = "Lobby"

func (c *conn) Create(name, passwd string) error {
	pHash, err := password.ComputeHash(passwd)
	if err != nil {
		return err
	}

	sChannel := storedChannel{
		Id:      bson.NewObjectId(),
		Name:    defaultChannelName,
		Created: time.Now(),
	}

	sRoom := storedRoom{
		Name:         name,
		PasswordHash: pHash,
		Channels:     []storedChannel{sChannel},
		Created:      time.Now(),
	}

	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err = col.Insert(&sRoom)
	if mgo.IsDup(err) {
		return db.DuplicateError
	}
	return err
}
