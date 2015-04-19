package accounts

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type conn struct {
	session *mgo.Session
	config  *config.Mongo
}

func New(session *mgo.Session, config *config.Mongo) (db.AccountsConn, error) {
	err := ensureIndices(session, config)
	return &conn{session, config}, err
}

func ensureIndices(session *mgo.Session, config *config.Mongo) error {
	emailIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     false,
	}
	col := session.DB(config.Name).C(config.Collections.Accounts)
	return col.EnsureIndex(emailIndex)
}
