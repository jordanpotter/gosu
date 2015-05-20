package rooms

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/api/v0/rooms/channels"
	"github.com/jordanpotter/gosu/server/api/v0/rooms/members"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type Handler struct {
	dbConn          *db.Conn
	tf              *token.Factory
	membersHandler  *members.Handler
	channelsHandler *channels.Handler
}

func New(dbConn *db.Conn, tf *token.Factory) *Handler {
	return &Handler{
		dbConn:          dbConn,
		tf:              tf,
		membersHandler:  members.New(dbConn),
		channelsHandler: channels.New(dbConn),
	}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.AuthRequired(h.tf))

	rg.POST("/", h.create)
	rg.GET("/id", h.getID)
	rg.GET("/id/:roomID", h.get)

	h.membersHandler.AddRoutes(rg.Group("/id/:roomID/members"))
	h.channelsHandler.AddRoutes(rg.Group("/id/:roomID/channels"))
}
