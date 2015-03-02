package mongo

import (
	"strings"

	mgo "gopkg.in/mgo.v2"

	"github.com/JordanPotter/gosu-server/internal/db"
)

const (
	databaseName = "gosu"
)

type conn struct {
	session *mgo.Session
}

func New(addresses []string) (db.Conn, error) {
	// TODO: use a username and password
	session, err := mgo.Dial(strings.Join(addresses, ","))
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Strong, false)
	session.SetSafe(&mgo.Safe{WMode: "majority", WTimeout: 1000, J: true})

	return &conn{session}, nil
}

func (c *conn) Close() {
	c.session.Close()
}
