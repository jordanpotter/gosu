package channels

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func (h *Handler) delete(c *gin.Context) {
	roomID := c.Params.ByName("roomID")
	channelID := c.Params.ByName("channelID")
	err := h.dbConn.Rooms.RemoveChannel(roomID, channelID)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
