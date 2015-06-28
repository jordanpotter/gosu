package pubsub

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/events"
)

type Subscriber interface {
	AddListener(listener chan<- SubMessage)
	SetAddrs(addrs []string) error
	Close() error
}

type SubMessage struct {
	Event     events.Event
	Timestamp time.Time
	Err       error
}
