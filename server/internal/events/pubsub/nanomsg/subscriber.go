package nanomsg

import (
	"fmt"
	"sync"

	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/sub"
	"github.com/gdamore/mangos/transport/tcp"

	"github.com/jordanpotter/gosu/server/internal/events"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
)

type subscriber struct {
	sock               mangos.Socket
	dialers            map[string]mangos.Dialer
	listeners          []chan<- pubsub.SubMessage
	listenersLock      sync.RWMutex
	messageHandlerOnce sync.Once
}

func NewSubscriber() (pubsub.Subscriber, error) {
	sock, err := sub.NewSocket()
	if err != nil {
		return nil, err
	}

	sock.AddTransport(tcp.NewTransport())
	return &subscriber{
		sock:    sock,
		dialers: make(map[string]mangos.Dialer),
	}, nil
}

func (s *subscriber) SetAddrs(addrs []string) error {
	err := s.addMissingConnections(addrs)
	if err != nil {
		return err
	}

	err = s.removeOldConnections(addrs)
	return err
}

func (s *subscriber) addMissingConnections(newAddrs []string) error {
	for _, addr := range newAddrs {
		_, exists := s.dialers[addr]
		if exists {
			continue
		}

		err := s.connect(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *subscriber) removeOldConnections(newAddrs []string) error {
Loop:
	for addr := range s.dialers {
		for _, newAddr := range newAddrs {
			if addr == newAddr {
				continue Loop
			}
		}

		err := s.disconnect(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *subscriber) connect(addr string) error {
	_, exists := s.dialers[addr]
	if exists {
		return fmt.Errorf("already connected to %s", addr)
	}

	dialer, err := s.sock.NewDialer("tcp://"+addr, nil)
	if err != nil {
		return err
	}

	err = dialer.Dial()
	if err != nil {
		return err
	}

	err = s.sock.SetOption(mangos.OptionSubscribe, []byte(""))
	if err != nil {
		dialer.Close()
		return err
	}

	s.dialers[addr] = dialer
	return nil
}

func (s *subscriber) disconnect(addr string) error {
	fmt.Println("disconnecting", addr)

	dialer, exists := s.dialers[addr]
	if !exists {
		return fmt.Errorf("not connected to %s", addr)
	}

	err := dialer.Close()
	if err != nil {
		return err
	}

	delete(s.dialers, addr)
	return nil
}

func (s *subscriber) AddListener(listener chan<- pubsub.SubMessage) {
	s.listenersLock.Lock()
	defer s.listenersLock.Unlock()

	s.listeners = append(s.listeners, listener)
	s.messageHandlerOnce.Do(s.handleMessages)
}

func (s *subscriber) handleMessages() {
	go func() {
		for {
			m := s.getNextMessage()
			s.listenersLock.RLock()
			for _, listener := range s.listeners {
				listener <- m
			}
			s.listenersLock.RUnlock()
		}
	}()
}

func (s *subscriber) getNextMessage() pubsub.SubMessage {
	b, err := s.sock.Recv()
	if err != nil {
		return pubsub.SubMessage{Err: err}
	}

	var m message
	err = msgpack.Unmarshal(b, &m)
	if err != nil {
		return pubsub.SubMessage{Err: err}
	}

	event, err := events.UnmarshalMsgpack(m.Type, m.EventData)
	if err != nil {
		return pubsub.SubMessage{Err: err}
	}

	return pubsub.SubMessage{
		Event:     event,
		Timestamp: m.Timestamp,
	}
}

func (s *subscriber) Close() error {
	err := s.sock.Close()
	s.dialers = make(map[string]mangos.Dialer)
	for _, listener := range s.listeners {
		close(listener)
	}
	return err
}
