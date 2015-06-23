package hub

import (
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
)

type Hub struct {
	clientManagers     map[string]*clientsManager
	clientManagersLock sync.Mutex
}

type sub struct {
	id      int
	f       http.Flusher
	created time.Time
}

func New(tf *token.Factory, sub pubsub.Subscriber) *Hub {
	return &Hub{
		clientManagers: make(map[string]*clientsManager),
	}
}

func (h *Hub) Send(key string, m []byte) {
	h.clientManagers[key].send(m)
}

func (h *Hub) SubscribeAndBlock(key string, w http.ResponseWriter) error {
	_, ok := w.(http.Flusher)
	if !ok {
		return errors.New("streaming not supported")
	}

	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Connection", "keep-alive")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	h.createClientManagerIfMissing(key)
	id := h.clientManagers[key].add(w)
	h.blockOnConnClose(w)
	return h.clientManagers[key].remove(id)
}

func (h *Hub) createClientManagerIfMissing(key string) {
	_, ok := h.clientManagers[key]
	if ok {
		return
	}

	h.clientManagersLock.Lock()
	defer h.clientManagersLock.Unlock()

	_, ok = h.clientManagers[key]
	if !ok {
		h.clientManagers[key] = new(clientsManager)
	}
}

func (h *Hub) blockOnConnClose(w http.ResponseWriter) {
	notify := w.(http.CloseNotifier).CloseNotify()
	<-notify
}
