package channels

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) delete(c *gin.Context) {
	roomName := c.Params.ByName("roomName")
	channelName := c.Params.ByName("channelName")

	fmt.Println("TODO: make sure admin for room")

	err := h.dbConn.Rooms.RemoveChannel(roomName, channelName)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
