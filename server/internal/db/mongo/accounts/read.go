package accounts

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) Get(id string) (*db.Account, error) {
	var sa storedAccount
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	err := col.FindId(bson.ObjectIdHex(id)).One(&sa)
	return handleStoredAccount(&sa, err)
}

func (c *conn) GetByEmail(email string) (*db.Account, error) {
	var sa storedAccount
	query := bson.M{"email": email}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	err := col.Find(query).One(&sa)
	return handleStoredAccount(&sa, err)
}

func handleStoredAccount(sa *storedAccount, err error) (*db.Account, error) {
	if err == mgo.ErrNotFound {
		return nil, db.NotFoundError
	} else if err != nil {
		return nil, err
	}
	return sa.toAccount(), nil
}
