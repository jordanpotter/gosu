package etcd

import (
	"encoding/json"

	"github.com/jordanpotter/gosu/server/internal/config"
)

const authTokenKey = "/auth/token"

func (c *conn) GetAuthToken() (*config.AuthToken, error) {
	resp, err := c.client.Get(authTokenKey, false, false)
	if err != nil {
		return nil, err
	}

	authTokenConfig := new(config.AuthToken)
	err = json.Unmarshal([]byte(resp.Node.Value), authTokenConfig)
	return authTokenConfig, err
}
