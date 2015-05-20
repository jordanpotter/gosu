package members

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type Handler struct {
	dbConn *db.Conn
}

func New(dbConn *db.Conn) *Handler {
	return &Handler{dbConn}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.GET("/", h.getAll)
	rg.POST("/join", h.join)
	rg.DELETE("/leave", h.leave)

	rg.PUT("/id/:memberID/admin", h.setAdmin)
	rg.PUT("/id/:memberID/banned", h.setBanned)
	rg.DELETE("/id/:memberID", h.delete)
}
