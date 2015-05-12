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
	rg.POST("/", h.join)
	rg.DELETE("/", h.leave)
	rg.PUT("/:memberName/admin", h.setAdmin)
	rg.PUT("/:memberName/banned", h.setBanned)
	rg.DELETE("/:memberName", h.delete)
}
