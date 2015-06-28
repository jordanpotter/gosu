package accounts

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/events/v0/sanitization"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/events"
	"github.com/jordanpotter/gosu/server/internal/events/hub"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

const listenerBufferSize = 100

type Handler struct {
	tf       *token.Factory
	sub      pubsub.Subscriber
	listener chan pubsub.SubMessage
	hub      *hub.Hub
}

func New(tf *token.Factory, sub pubsub.Subscriber) *Handler {
	h := &Handler{
		tf:       tf,
		sub:      sub,
		listener: make(chan pubsub.SubMessage, listenerBufferSize),
		hub:      hub.New(),
	}
	h.sub.AddListener(h.listener)
	go h.handleEvents()
	return h
}

func (h *Handler) handleEvents() {
	for sm := range h.listener {
		if sm.Err != nil {
			fmt.Println("Received subscription error", sm.Err)
			continue
		}

		var key string
		var sanitizedEvent interface{}
		switch e := sm.Event.(type) {
		case events.AccountDeviceCreated:
			key = strconv.Itoa(e.AccountID)
			sanitizedEvent = sanitization.ToAccountDeviceCreated(e, sm.Timestamp)
		case events.AccountDeviceDeleted:
			key = strconv.Itoa(e.AccountID)
			sanitizedEvent = sanitization.ToAccountDeviceDeleted(e, sm.Timestamp)
		default:
			continue
		}

		b, err := json.Marshal(sanitizedEvent)
		if err != nil {
			fmt.Println("Error while processing event", err)
			continue
		}

		h.hub.Send(key, b)
	}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.AuthRequired(h.tf))

	rgWithID := rg.Group("/id/:accountID")
	rgWithID.Use(middleware.AuthMatchesAccount("accountID"))
	rgWithID.POST("/subscribe", h.subscribe)
}

func (h *Handler) subscribe(c *gin.Context) {
	t, ok := c.Get(middleware.TokenKey)
	if !ok {
		c.AbortWithError(500, errors.New("missing auth token"))
		return
	}
	authToken := t.(*token.Token)

	err := h.hub.SubscribeAndBlock(strconv.Itoa(authToken.Account.ID), c.Writer.(http.ResponseWriter))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.String(200, "ok")
}
