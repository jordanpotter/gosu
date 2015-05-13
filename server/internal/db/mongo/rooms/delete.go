package rooms

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func (c *conn) Delete(name string) error {
	query := bson.M{"name": name}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	err := col.Remove(query)
	if err == mgo.ErrNotFound {
		return db.NotFoundError
	}
	return err
}
