package etcd

import (
	"time"

	"github.com/coreos/go-etcd/etcd"
	"github.com/jordanpotter/gosu/server/internal/config"
)

type conn struct {
	client *etcd.Client
}

func New(addrs []string) config.Conn {
	client := etcd.NewClient(addrs)
	client.SetDialTimeout(10 * time.Second)
	client.SetConsistency("STRONG_CONSISTENCY")
	return &conn{client}
}

func (c *conn) Close() {
	c.Close()
}
