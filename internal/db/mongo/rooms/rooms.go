package rooms

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/JordanPotter/gosu-server/internal/config"
	"github.com/JordanPotter/gosu-server/internal/db"
)

type conn struct {
	session *mgo.Session
	config  *config.DBMongo
}

func New(session *mgo.Session, config *config.DBMongo) (db.RoomsConn, error) {
	err := ensureIndices(session, config)
	if err != nil {
		return nil, err
	}
	return &conn{session, config}, nil
}

func ensureIndices(session *mgo.Session, config *config.DBMongo) error {
	nameIndex := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     false,
	}
	col := session.DB(config.Name).C(config.Collections.Rooms)
	return col.EnsureIndex(nameIndex)
}
