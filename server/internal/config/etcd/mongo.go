package etcd

import (
	"encoding/json"

	"github.com/jordanpotter/gosu/server/internal/config"
)

const mongoConfKey = "/conf/mongo"

func (c *conn) GetMongo() (*config.Mongo, error) {
	resp, err := c.client.Get(mongoConfKey, false, false)
	if err != nil {
		return nil, err
	}

	mongoConfig := new(config.Mongo)
	err = json.Unmarshal([]byte(resp.Node.Value), mongoConfig)
	return mongoConfig, err
}
