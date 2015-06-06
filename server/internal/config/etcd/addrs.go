package etcd

import (
	"encoding/json"

	"github.com/jordanpotter/gosu/server/internal/config"
)

const (
	addrsKey         = "/addrs"
	addrsAuthKey     = "/addrs/auth"
	addrsAPIKey      = "/addrs/api"
	addrsEventsKey   = "/addrs/events"
	addrsRelayKey    = "/addrs/relay"
	addrsPostgresKey = "/addrs/postgres"
)

func (c *conn) GetAuthAddrs() ([]config.AuthNode, error) {
	resp, err := c.client.Get(addrsAuthKey, true, false)
	if err != nil {
		return nil, err
	}

	addrs := make([]config.AuthNode, 0, len(resp.Node.Nodes))
	for _, node := range resp.Node.Nodes {
		authNode := config.AuthNode{}
		err = json.Unmarshal([]byte(node.Value), &authNode)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, authNode)
	}
	return addrs, nil
}

func (c *conn) GetAPIAddrs() ([]config.APINode, error) {
	resp, err := c.client.Get(addrsAPIKey, true, false)
	if err != nil {
		return nil, err
	}

	addrs := make([]config.APINode, 0, len(resp.Node.Nodes))
	for _, node := range resp.Node.Nodes {
		apiNode := config.APINode{}
		err = json.Unmarshal([]byte(node.Value), &apiNode)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, apiNode)
	}
	return addrs, nil
}

func (c *conn) GetEventsAddrs() ([]config.EventsNode, error) {
	resp, err := c.client.Get(addrsEventsKey, true, false)
	if err != nil {
		return nil, err
	}

	addrs := make([]config.EventsNode, 0, len(resp.Node.Nodes))
	for _, node := range resp.Node.Nodes {
		eventsNode := config.EventsNode{}
		err = json.Unmarshal([]byte(node.Value), &eventsNode)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, eventsNode)
	}
	return addrs, nil
}

func (c *conn) GetRelayAddrs() ([]config.RelayNode, error) {
	resp, err := c.client.Get(addrsRelayKey, true, false)
	if err != nil {
		return nil, err
	}

	addrs := make([]config.RelayNode, 0, len(resp.Node.Nodes))
	for _, node := range resp.Node.Nodes {
		relayNode := config.RelayNode{}
		err = json.Unmarshal([]byte(node.Value), &relayNode)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, relayNode)
	}
	return addrs, nil
}

func (c *conn) GetPostgresAddrs() ([]config.PostgresNode, error) {
	resp, err := c.client.Get(addrsPostgresKey, true, false)
	if err != nil {
		return nil, err
	}

	addrs := make([]config.PostgresNode, 0, len(resp.Node.Nodes))
	for _, node := range resp.Node.Nodes {
		postgresNode := config.PostgresNode{}
		err = json.Unmarshal([]byte(node.Value), &postgresNode)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, postgresNode)
	}
	return addrs, nil
}
