package rooms

import (
	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu-server/internal/auth/token"
	"github.com/JordanPotter/gosu-server/internal/db"
)

type Handler struct {
	dbConn       db.Conn
	tokenFactory *token.Factory
}

func New(dbConn db.Conn, tokenFactory *token.Factory) *Handler {
	return &Handler{dbConn, tokenFactory}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/", h.create)
	rg.GET("/:name", h.get)
	rg.POST("/:name/login", h.login)
	rg.POST("/:name/logout", h.logout)

	rg.PUT("/:name/password", h.setPassword)

	rg.GET("/:name/users", h.getUsers)
	rg.PUT("/:name/users/:userName/admin", h.setAdmin)
	rg.PUT("/:name/users/:userName/banned", h.setBanned)
	rg.DELETE("/:name/users/:userName", h.deleteUser)

	rg.POST("/:name/channels", h.createChannel)
	rg.DELETE("/:name/channels/:channelName", h.deleteChannel)
	rg.POST("/:name/channels/:channelName/move", h.moveChannel)
	rg.GET("/:name/channels/:channelName/relays", h.getRelayConns)
}
