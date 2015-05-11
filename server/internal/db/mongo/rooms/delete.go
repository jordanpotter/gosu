package rooms

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) Delete(name string) error {
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err := col.Remove(bson.M{"name": name})
	if err == mgo.ErrNotFound {
		return db.NotFoundError
	}
	return err
}
