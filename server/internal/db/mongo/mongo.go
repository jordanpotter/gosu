package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/db/mongo/accounts"
	"github.com/jordanpotter/gosu/server/internal/db/mongo/rooms"
)

func New(config *config.DBMongo) (*db.Conn, error) {
	session, err := createSession(config)
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

func createSession(config *config.DBMongo) (*mgo.Session, error) {
	dialInfo := mgo.DialInfo{
		Addrs:     config.Addresses,
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
	return session, nil
}
