package accounts

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JordanPotter/gosu-server/server/internal/db"
)

func (c *conn) Get(email string) (*db.Account, error) {
	var sa storedAccount
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	err := col.Find(bson.M{"email": email}).One(&sa)
	if err == mgo.ErrNotFound {
		return nil, db.NotFoundError
	} else if err != nil {
		return nil, err
	}
	return sa.toAccount(), nil
}
