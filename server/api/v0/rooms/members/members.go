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
	rg.POST("/authenticate", h.authenticate)
	rg.DELETE("/leave", h.leave)

	rg.PUT("/manage/:memberAccountID/admin", h.setAdmin)
	rg.PUT("/manage/:memberAccountID/banned", h.setBanned)
	rg.DELETE("/manage/:memberAccountID", h.delete)
}
