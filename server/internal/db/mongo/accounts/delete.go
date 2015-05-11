package accounts

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) Delete(id string) error {
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	err := col.RemoveId(bson.ObjectIdHex(id))
	if err == mgo.ErrNotFound {
		return db.NotFoundError
	}
	return err
}
