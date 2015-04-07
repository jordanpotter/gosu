package accounts

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/JordanPotter/gosu/server/internal/config"
	"github.com/JordanPotter/gosu/server/internal/db"
)

type conn struct {
	session *mgo.Session
	config  *config.DBMongo
}

func New(session *mgo.Session, config *config.DBMongo) (db.AccountsConn, error) {
	err := ensureIndices(session, config)
	return &conn{session, config}, err
}

func ensureIndices(session *mgo.Session, config *config.DBMongo) error {
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
