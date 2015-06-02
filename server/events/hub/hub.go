package hub

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/middleware"
	"github.com/jordanpotter/gosu/server/internal/pubsub"
)

const (
	subMessageBuffer = 100
)

type Hub struct {
	tf             *token.Factory
	listenChan     chan *pubsub.SubMessage
	roomsToClients map[string]*clientsManager
}

type sub struct {
	id      int
	f       http.Flusher
	created time.Time
}

func New(tf *token.Factory, sub pubsub.Subscriber) (*Hub, error) {
	h := &Hub{
		tf:             tf,
		listenChan:     make(chan *pubsub.SubMessage, subMessageBuffer),
		roomsToClients: make(map[string]*clientsManager),
	}

	err := sub.Listen(h.listenChan)
	if err != nil {
		return nil, err
	}

	go h.handleMessages()
	return h, nil
}

func (h *Hub) handleMessages() {
	for m := range h.listenChan {
		fmt.Println(m.Event)
	}
}

func (h *Hub) AddRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.AuthRequired(h.tf), middleware.AuthMatchesRoom("roomID"))
	rg.POST(":roomID/subscribe", h.subscribe)
}

func (h *Hub) subscribe(c *gin.Context) {
	_, ok := c.Writer.(http.Flusher)
	if !ok {
		c.AbortWithError(500, errors.New("streaming not supported"))
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	roomID := c.Params.ByName("roomID")
	id := h.addToRoomClients(roomID, c.Writer)
	h.blockOnConnClose(c.Writer)
	err := h.roomsToClients[roomID].remove(id)
	if err != nil {
		panic(err)
	}

	c.String(200, "ok")
}

func (h *Hub) addToRoomClients(roomID string, w gin.ResponseWriter) int {
	_, ok := h.roomsToClients[roomID]
	if !ok {
		h.roomsToClients[roomID] = new(clientsManager)
	}
	return h.roomsToClients[roomID].add(w)
}

func (h *Hub) blockOnConnClose(w gin.ResponseWriter) {
	notify := w.(http.CloseNotifier).CloseNotify()
	<-notify
}
