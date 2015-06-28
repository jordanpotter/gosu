package hub

import (
	"fmt"
	"net/http"
	"sync"
)

type clientsManager struct {
	clients []client
	lock    sync.RWMutex
	counter int
}

type client struct {
	id int
	w  http.ResponseWriter
}

func (cm *clientsManager) add(w http.ResponseWriter) int {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	c := client{cm.counter, w}
	cm.counter++
	cm.clients = append(cm.clients, c)
	return c.id
}

func (cm *clientsManager) remove(id int) error {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	for i, c := range cm.clients {
		if c.id == id {
			cm.clients = append(cm.clients[:i], cm.clients[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("missing client with id %d", id)
}

func (cm *clientsManager) send(m []byte) {
	cm.lock.RLock()
	defer cm.lock.RUnlock()

	for _, c := range cm.clients {
		go func(c client) {
			fmt.Fprint(c.w, m)
			c.w.(http.Flusher).Flush()
		}(c)
	}
}
