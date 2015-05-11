package rooms

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/api/middleware"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func (h *Handler) get(c *gin.Context) {
	roomName, err := c.Get(middleware.RoomNameKey)
	if err != nil {
		c.Fail(500, err)
		return
	}

	room, err := h.dbConn.Rooms.Get(roomName.(string))
	if err == db.NotFoundError {
		c.Fail(404, fmt.Errorf("missing room %s", roomName.(string)))
		return
	}

	c.JSON(200, room)
}
