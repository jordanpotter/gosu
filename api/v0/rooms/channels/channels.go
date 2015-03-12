package channels

import (
	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu-server/internal/db"
)

type Handler struct {
	dbConn db.Conn
}

func New(dbConn db.Conn) *Handler {
	return &Handler{dbConn}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/", h.create)
	rg.DELETE("/:channelName", h.delete)
	rg.POST("/:channelName/move", h.move)
	rg.GET("/:channelName/relays", h.getRelayConns)
}
