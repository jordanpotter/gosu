package mongo

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/JordanPotter/gosu-server/internal/db"
)

type conn struct {
	session *mgo.Session
}

func New() (db.Conn, error) {
	s, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}

	// TODO: use a username and password
	// TODO: consistent writes as default
	return &conn{s}, nil
}

func (c *conn) CreateAccount(name string, password string) (*db.Account, error) {
	return nil, nil
}

func (c *conn) GetAccount(name string) (*db.Account, error) {
	return nil, nil
}

func (c *conn) Close() {
	c.session.Close()
}
