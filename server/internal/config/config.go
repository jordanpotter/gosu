package config

import (
	"encoding/hex"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Global struct {
	API    API    `yaml:"api"`
	Events Events `yaml:"events"`
	Relay  Relay  `yaml:"relay"`
	DB     DB     `yaml:"db"`
	Auth   Auth   `yaml:"auth"`
}

type API struct {
	Port int `yaml:"port"`
}

type Events struct {
	Port int `yaml:"port"`
}

type Relay struct {
	Port int `yaml:"port"`
}

type DB struct {
	Mongo DBMongo `yaml:"mongo"`
}

type DBMongo struct {
	Addresses    []string           `yaml:"addresses"`
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

func Load(path string) (*Global, error) {
	file, err := ioutil.ReadFile(path)
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
