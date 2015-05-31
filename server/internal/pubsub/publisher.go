package pubsub

import "github.com/jordanpotter/gosu/server/events/types"

type Publisher interface {
	Send(event types.Event) error
	Close() error
}
