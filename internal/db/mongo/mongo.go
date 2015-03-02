package mongo

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JordanPotter/gosu-server/internal/db"
)

type conn struct {
	session *mgo.Session
}

func New() (db.Conn, error) {
	// TODO: use a username and password
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Strong, false)
	session.SetSafe(&mgo.Safe{WMode: "majority", WTimeout: 1000, J: true})

	ensureIndices(session)

	return &conn{session}, nil
}

func ensureIndices(session *mgo.Session) error {
	nameIndex := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}
	return session.DB("gosu").C("accounts").EnsureIndex(nameIndex)
}

func (c *conn) CreateAccount(name, password string) (*db.Account, error) {
	err := db.CheckAccountName(name)
	if err != nil {
		return nil, err
	}

	pHash, err := db.ComputePasswordHash(password)
	if err != nil {
		return nil, err
	}

	account := &db.Account{
		Name:         name,
		PasswordHash: pHash,
	}

	col := c.session.DB("gosu").C("accounts")
	err = col.Insert(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (c *conn) GetAccount(name, password string) (*db.Account, error) {
	var account db.Account
	col := c.session.DB("gosu").C("accounts")
	err := col.Find(bson.M{"name": name}).One(&account)
	if err != nil {
		return nil, err
	}

	if !db.DoesPasswordMatchHash(password, account.PasswordHash) {
		return nil, errors.New("mongo: password does not match hash")
	}

	return &account, nil
}

func (c *conn) Close() {
	c.session.Close()
}
