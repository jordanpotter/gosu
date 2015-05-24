package etcd

const (
	addrsKey       = "/addrs"
	addrsAuthKey   = "/addrs/auth"
	addrsAPIKey    = "/addrs/api"
	addrsEventsKey = "/addrs/events"
	addrsRelayKey  = "/addrs/relay"
	addrsMongoKey  = "/addrs/mongo"
)

func (c *conn) GetAuthAddrs() ([]string, error) {
	return c.getAddrs(addrsAuthKey)
}

func (c *conn) GetAPIAddrs() ([]string, error) {
	return c.getAddrs(addrsAPIKey)
}

func (c *conn) GetEventsAddrs() ([]string, error) {
	return c.getAddrs(addrsEventsKey)
}

func (c *conn) GetRelayAddrs() ([]string, error) {
	return c.getAddrs(addrsRelayKey)
}

func (c *conn) GetMongoAddrs() ([]string, error) {
	return c.getAddrs(addrsMongoKey)
}

func (c *conn) getAddrs(key string) ([]string, error) {
	resp, err := c.client.Get(key, true, false)
	if err != nil {
		return nil, err
	}

	addrs := make([]string, 0, len(resp.Node.Nodes))
	for _, node := range resp.Node.Nodes {
		addrs = append(addrs, node.Value)
	}
	return addrs, nil
}
