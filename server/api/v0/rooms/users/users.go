package users

import (
	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu/server/internal/db"
)

type Handler struct {
	dbConn *db.Conn
}

func New(dbConn *db.Conn) *Handler {
	return &Handler{dbConn}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.GET("/", h.getAll)
	rg.PUT("/:userName/admin", h.setAdmin)
	rg.PUT("/:userName/banned", h.setBanned)
	rg.DELETE("/:userName", h.delete)
}
