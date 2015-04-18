package rooms

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/api/v0/rooms/channels"
	"github.com/jordanpotter/gosu/server/api/v0/rooms/users"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type Handler struct {
	dbConn          *db.Conn
	tokenFactory    *token.Factory
	usersHandler    *users.Handler
	channelsHandler *channels.Handler
}

func New(dbConn *db.Conn, tokenFactory *token.Factory) *Handler {
	return &Handler{
		dbConn:          dbConn,
		tokenFactory:    tokenFactory,
		usersHandler:    users.New(dbConn),
		channelsHandler: channels.New(dbConn),
	}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/", h.create)
	rg.GET("/:roomName", h.get)
	rg.POST("/:roomName/login", h.login)
	rg.POST("/:roomName/logout", h.logout)
	rg.PUT("/:roomName/password", h.setPassword)

	h.usersHandler.AddRoutes(rg.Group("/:roomName/users"))
	h.channelsHandler.AddRoutes(rg.Group("/:roomName/channels"))
}
