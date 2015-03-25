package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/JordanPotter/gosu-server/internal/config"
	"github.com/JordanPotter/gosu-server/internal/db"
)

type conn struct {
	session *mgo.Session
	config  *config.DBMongo
}

func New(config *config.DBMongo) (db.Conn, error) {
	dialInfo := mgo.DialInfo{
		Addrs:     []string{config.Address},
		Database:  config.Name,
		Username:  config.Username,
		Password:  config.Password,
		Mechanism: "SCRAM-SHA-1",
		Direct:    false,
		Timeout:   10 * time.Second,
	}
	session, err := mgo.DialWithInfo(&dialInfo)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Strong, false)
	session.SetSafe(&mgo.Safe{WMode: config.WriteMode, WTimeout: config.WriteTimeout, J: config.Journaling})

	return &conn{session, config}, nil
}

func (c *conn) Close() {
	c.session.Close()
}
