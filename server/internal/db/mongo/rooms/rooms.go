package rooms

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type conn struct {
	session *mgo.Session
	config  *config.Mongo
}

func New(session *mgo.Session, config *config.Mongo) (db.RoomsConn, error) {
	err := ensureIndices(session, config)
	return &conn{session, config}, err
}

func ensureIndices(session *mgo.Session, config *config.Mongo) error {
	col := session.DB(config.Name).C(config.Collections.Rooms)

	nameIndex := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     false,
	}
	return col.EnsureIndex(nameIndex)
}
