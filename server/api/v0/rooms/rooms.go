package rooms

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/api/v0/rooms/channels"
	"github.com/jordanpotter/gosu/server/api/v0/rooms/members"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/middleware"
	"github.com/jordanpotter/gosu/server/internal/pubsub"
)

type Handler struct {
	dbConn          db.Conn
	tf              *token.Factory
	pub             pubsub.Publisher
	membersHandler  *members.Handler
	channelsHandler *channels.Handler
}

func New(dbConn db.Conn, tf *token.Factory, pub pubsub.Publisher) *Handler {
	return &Handler{
		dbConn:          dbConn,
		tf:              tf,
		pub:             pub,
		membersHandler:  members.New(dbConn, tf, pub),
		channelsHandler: channels.New(dbConn, tf, pub),
	}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.AuthRequired(h.tf))
	rg.POST("/", h.create)
	rg.GET("/id", h.getID)

	rgWithID := rg.Group("/id/:roomID")
	rgWithID.Use(middleware.AuthMatchesRoom("roomID"))
	rgWithID.GET("", h.get)
	h.membersHandler.AddRoutes(rgWithID.Group("/members"))
	h.channelsHandler.AddRoutes(rgWithID.Group("/channels"))
}
