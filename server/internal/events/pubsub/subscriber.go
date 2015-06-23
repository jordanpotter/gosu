package pubsub

import "time"

type Subscriber interface {
	Listen(listener chan<- *SubMessage) error
	SetAddrs(addrs []string) error
	Close() error
}

type SubMessage struct {
	Name      string
	Data      []byte
	Timestamp time.Time
	Err       error
}
