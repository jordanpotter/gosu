package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/db/mongo/accounts"
	"github.com/jordanpotter/gosu/server/internal/db/mongo/rooms"
)

func New(addrs []string, config *config.Mongo) (*db.Conn, error) {
	session, err := createSession(addrs, config)
	if err != nil {
		return nil, err
	}

	accountsConn, err := accounts.New(session, config)
	if err != nil {
		return nil, err
	}

	roomsConn, err := rooms.New(session, config)
	if err != nil {
		return nil, err
	}

	conn := &db.Conn{
		Accounts: accountsConn,
		Rooms:    roomsConn,
		Closer:   session,
	}
	return conn, nil
}

func createSession(addrs []string, config *config.Mongo) (*mgo.Session, error) {
	dialInfo := mgo.DialInfo{
		Addrs:     addrs,
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
	session.SetSafe(&mgo.Safe{
		WMode:    config.WriteParams.Mode,
		WTimeout: int(config.WriteParams.Timeout.Seconds()),
		J:        config.WriteParams.Journaling})
	return session, nil
}
