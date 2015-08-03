package nanomsg

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/jordanpotter/gosu/server/internal/events"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
)

func createPublisher(t *testing.T) (pubsub.Publisher, string) {
	testPort := rand.Intn(int(math.Exp2(16)-math.Exp2(10))) + int(math.Exp2(10))
	addr := fmt.Sprintf("127.0.0.1:%d", testPort)

	pub, err := NewPublisher(addr)
	if err != nil {
		t.Fatalf("Unexpected error while creating publisher: %v", err)
	}
	return pub, addr
}

func createSubscriber(t *testing.T, addr string) pubsub.Subscriber {
	sub, err := NewSubscriber()
	if err != nil {
		t.Fatalf("Unexpected error while creating subscriber: %v", err)
	}

	err = sub.SetAddrs([]string{addr})
	if err != nil {
		t.Fatalf("Unexpected error while setting subscriber addresses: %v", err)
	}
	return sub
}

func createEvent(id int) events.Event {
	return events.RoomChannelCreated{
		RoomID:      id,
		ChannelID:   id,
		ChannelName: fmt.Sprintf("channel-%d", id),
		Created:     time.Now(),
	}
}

func TestSending(t *testing.T) {
	pub, addr := createPublisher(t)
	sub := createSubscriber(t, addr)
	listener := make(chan pubsub.SubMessage, 100)
	sub.AddListener(listener)

	for i := 0; i < 100; i++ {
		event := createEvent(i)
		pub.Send(event)

		subEvent := <-listener
		if subEvent.Err != nil {
			t.Fatalf("Unexpected error: %v", subEvent.Err)
		} else if !reflect.DeepEqual(event, subEvent.Event) {
			t.Errorf("Events are not equal: %v != %v", event, subEvent.Event)
		}
	}
}

func TestHeavyLoad(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	numEvents := 100000
	pub, addr := createPublisher(t)
	sub := createSubscriber(t, addr)
	listener := make(chan pubsub.SubMessage, numEvents)
	sub.AddListener(listener)

	go func() {
		for i := 0; i < numEvents; i++ {
			time.Sleep(10 * time.Microsecond)
			event := createEvent(i)
			pub.Send(event)
		}
	}()

	for i := 0; i < numEvents; i++ {
		subEvent := <-listener
		if subEvent.Err != nil {
			t.Fatalf("Unexpected error: %v", subEvent.Err)
		}
	}
}
