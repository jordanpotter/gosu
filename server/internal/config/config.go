package config

import "time"

type Conn interface {
	GetAuthToken() (*AuthToken, error)
	GetMongo() (*Mongo, error)

	GetAuthAddrs() ([]string, error)
	GetAPIAddrs() ([]string, error)
	GetEventsAddrs() ([]string, error)
	GetRelayAddrs() ([]string, error)
	GetMongoAddrs() ([]string, error)

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
