package channels

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/events/types"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func (h *Handler) delete(c *gin.Context) {
	roomID := c.Params.ByName("roomID")
	channelID := c.Params.ByName("channelID")
	err := h.dbConn.Rooms.RemoveChannel(roomID, channelID)
	if err == db.NotFoundError {
		c.AbortWithError(404, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	e := &types.RoomChannelDeleted{
		RoomID:    roomID,
		ChannelID: channelID,
	}
	err = h.pub.Send(e)
	if err != nil {
		fmt.Printf("Failed to send event: %v", err)
	}

	c.String(200, "ok")
}
