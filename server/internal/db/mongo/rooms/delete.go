package rooms

import "gopkg.in/mgo.v2/bson"

func (c *conn) Delete(id string) error {
	col := c.session.DB(c.config.Name).C(c.config.Collections.Rooms)
	return col.RemoveId(bson.ObjectIdHex(id))
}
