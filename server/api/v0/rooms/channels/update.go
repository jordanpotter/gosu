package channels

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) move(c *gin.Context) {
	roomName := c.Params.ByName("roomName")
	channelName := c.Params.ByName("channelName")
	fmt.Printf("TODO: move to channel %s for room %s\n", channelName, roomName)

	c.String(200, "ok")
}
