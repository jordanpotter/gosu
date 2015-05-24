package events

type Subscriber interface {
	Listen(listener chan<- *Message) error
	SetAddrs(addrs []string) error
	Close() error
}
