package channels

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/events"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type Handler struct {
	dbConn *db.Conn
	pub    events.Publisher
}

func New(dbConn *db.Conn, pub events.Publisher) *Handler {
	return &Handler{dbConn, pub}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.IsRoomAdmin())
	rg.POST("/", h.create)
	rg.DELETE("/id/:channelID", h.delete)
}
