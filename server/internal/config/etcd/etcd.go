package etcd

import (
	"time"

	"github.com/coreos/go-etcd/etcd"
	"github.com/jordanpotter/gosu/server/internal/config"
)

type conn struct {
	client *etcd.Client
}

func New(addresses []string) config.Conn {
	client := etcd.NewClient(addresses)
	client.SetDialTimeout(10 * time.Second)
	client.SetConsistency("STRONG_CONSISTENCY")
	return &conn{client}
}
