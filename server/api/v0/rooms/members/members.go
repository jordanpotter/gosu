package members

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
	rg.POST("/join", h.join)
	rg.DELETE("/leave", h.leave)

	rgWithID := rg.Group("/id/:memberID")
	rgWithID.Use(middleware.IsRoomAdmin())
	rgWithID.PUT("/id/:memberID/admin", h.setAdmin)
	rgWithID.PUT("/id/:memberID/banned", h.setBanned)
	rgWithID.DELETE("/id/:memberID", h.delete)
}
