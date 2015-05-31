package pubsub

import (
	"time"

	"github.com/jordanpotter/gosu/server/events/types"
)

type Subscriber interface {
	Listen(listener chan<- *SubMessage) error
	SetAddrs(addrs []string) error
	Close() error
}

type SubMessage struct {
	Event     types.Event
	Timestamp time.Time
	Err       error
}
