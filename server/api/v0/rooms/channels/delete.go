package channels

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func (h *Handler) delete(c *gin.Context) {
	roomIDString := c.Params.ByName("roomID")
	roomID, err := strconv.Atoi(roomIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	channelIDString := c.Params.ByName("channelID")
	channelID, err := strconv.Atoi(channelIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = h.dbConn.DeleteChannelForRoom(channelID, roomID)
	if err == db.NotFoundError {
		c.AbortWithError(404, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// e := &types.RoomChannelDeleted{
	// 	RoomID:    roomID,
	// 	ChannelID: channelID,
	// }
	// err = h.pub.Send(e)
	// if err != nil {
	// 	fmt.Printf("Failed to send event: %v", err)
	// }

	c.String(200, "ok")
}
