package pubsub

import "github.com/jordanpotter/gosu/server/internal/events"

type Publisher interface {
	Send(event events.Event) error
	Close() error
}
