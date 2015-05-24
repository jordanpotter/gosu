package events

type Subscriber interface {
	Listen() <-chan Message
	Close() error
}
