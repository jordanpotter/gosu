package config

import (
	"net"
	"time"
)

type Conn interface {
	GetAuthToken() (*AuthToken, error)
	GetPostgres() (*Postgres, error)

	GetAuthAddrs() ([]AuthNode, error)
	GetAPIAddrs() ([]APINode, error)
	GetEventsAddrs() ([]EventsNode, error)
	GetRelayAddrs() ([]RelayNode, error)
	GetPostgresAddrs() ([]PostgresNode, error)

	Close()
}

type AuthToken struct {
	Key      []byte        `json:"key"`
	Duration time.Duration `json:"durationNs"`
}

type Postgres struct {
	Username string         `json:"username"`
	Password string         `json:"password"`
	Name     string         `json:"name"`
	Tables   PostgresTables `json:"tables"`
	SSLMode  string         `json:"sslMode"`
}

type PostgresTables struct {
	Accounts string `json:"accounts"`
	Devices  string `json:"devices"`
	Rooms    string `json:"rooms"`
	Channels string `json:"channels"`
	Members  string `json:"members"`
}

type AuthNode struct {
	IP       net.IP `json:"ip"`
	HttpPort int    `json:"httpPort"`
}

type APINode struct {
	IP       net.IP `json:"ip"`
	HttpPort int    `json:"httpPort"`
	PubPort  int    `json:"pubPort"`
}

type EventsNode struct {
	IP       net.IP `json:"ip"`
	HttpPort int    `json:"httpPort"`
	SubPort  int    `json:"subPort"`
}

type RelayNode struct {
	IP        net.IP `json:"ip"`
	HttpPort  int    `json:"httpPort"`
	CommsPort int    `json:"commsPort"`
}

type PostgresNode struct {
	IP     net.IP `json:"ip"`
	DBPort int    `json:"dbPort"`
}
