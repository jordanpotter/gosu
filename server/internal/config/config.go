package config

import (
	"net"
	"time"
)

type Conn interface {
	GetAuthToken() (*AuthToken, error)
	GetMongo() (*Mongo, error)

	GetAuthAddrs() ([]AuthNode, error)
	GetAPIAddrs() ([]APINode, error)
	GetEventsAddrs() ([]EventsNode, error)
	GetRelayAddrs() ([]RelayNode, error)
	GetMongoAddrs() ([]MongoNode, error)

	Close()
}

type AuthToken struct {
	Key      []byte        `json:"key"`
	Duration time.Duration `json:"durationNs"`
}

type Mongo struct {
	Username    string           `json:"username"`
	Password    string           `json:"password"`
	Name        string           `json:"name"`
	Collections MongoCollections `json:"collections"`
	WriteParams MongoWriteParams `json:"writeParams"`
}

type MongoCollections struct {
	Accounts string `json:"accounts"`
	Rooms    string `json:"rooms"`
}

type MongoWriteParams struct {
	Mode       string        `json:"mode"`
	Journaling bool          `json:"journaling"`
	Timeout    time.Duration `json:"timeoutNs"`
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

type MongoNode struct {
	IP     net.IP `json:"ip"`
	DBPort int    `json:"dbPort"`
}
