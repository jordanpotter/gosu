package accounts

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/events/hub"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type Handler struct {
	tf  *token.Factory
	sub pubsub.Subscriber
	hub *hub.Hub
}

func New(tf *token.Factory, sub pubsub.Subscriber) *Handler {
	return &Handler{
		tf:  tf,
		sub: sub,
		hub: hub.New(),
	}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.AuthRequired(h.tf))

	rgWithID := rg.Group("/id/:accountID")
	rgWithID.Use(middleware.AuthMatchesAccount("accountID"))
	rgWithID.POST("/subscribe", h.subscribe)
}

func (h *Handler) subscribe(c *gin.Context) {
	c.String(200, "ok")
}
