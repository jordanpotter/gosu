package config

import (
	"encoding/hex"
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Global struct {
	API    API    `yaml:"api"`
	Events Events `yaml:"events"`
	Relay  Relay  `yaml:"relay"`
	DB     DB     `yaml:"db"`
	Auth   Auth   `yaml:"auth"`
}

type API struct {
	Address string `yaml:"address"`
}

type Events struct {
	Address string `yaml:"address"`
}

type Relay struct {
	Address string `yaml:"address"`
}

type DB struct {
	Mongo DBMongo `yaml:"mongo"`
}

type DBMongo struct {
	Address      string             `yaml:"address"`
	Username     string             `yaml:"username"`
	Password     string             `yaml:"password"`
	Name         string             `yaml:"name"`
	Collections  DBMongoCollections `yaml:"collections"`
	WriteMode    string             `yaml:"writeMode"`
	WriteTimeout int                `yaml:"writeTimeout"`
	Journaling   bool               `yaml:"journaling"`
}

type DBMongoCollections struct {
	Accounts string `yaml:"accounts"`
	Rooms    string `yaml:"rooms"`
}

type Auth struct {
	Token AuthToken `yaml:"token"`
}

type AuthToken struct {
	Key      []byte        `yaml:"-"`
	KeyInHex string        `yaml:"keyInHex"`
	Duration time.Duration `yaml:"duration"`
}

func Get() (*Global, error) {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	globalConfig := new(Global)
	err = yaml.Unmarshal(file, globalConfig)
	if err != nil {
		return nil, err
	}

	err = populateRemainingFields(globalConfig)
	return globalConfig, err
}

func populateRemainingFields(global *Global) error {
	var err error
	global.Auth.Token.Key, err = hex.DecodeString(global.Auth.Token.KeyInHex)
	return err
}
