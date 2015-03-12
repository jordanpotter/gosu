package channels

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getRelayConns(c *gin.Context) {
	roomName := c.Params.ByName("roomName")
	channelName := c.Params.ByName("channelName")
	fmt.Printf("TODO: get relay connections for channel %s in room %s\n", channelName, roomName)

	c.String(200, "ok")
}
