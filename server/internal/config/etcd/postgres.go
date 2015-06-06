package etcd

import (
	"encoding/json"

	"github.com/jordanpotter/gosu/server/internal/config"
)

const postgresConfKey = "/conf/postgres"

func (c *conn) GetPostgres() (*config.Postgres, error) {
	resp, err := c.client.Get(postgresConfKey, false, false)
	if err != nil {
		return nil, err
	}

	postgresConfig := new(config.Postgres)
	err = json.Unmarshal([]byte(resp.Node.Value), postgresConfig)
	return postgresConfig, err
}
