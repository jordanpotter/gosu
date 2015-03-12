package rooms

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) deleteUser(c *gin.Context) {
	name := c.Params.ByName("name")
	userName := c.Params.ByName("userName")
	fmt.Printf("TODO: delete user %s in room %s\n", userName, name)

	c.String(200, "ok")
}

func (h *Handler) deleteChannel(c *gin.Context) {
	name := c.Params.ByName("name")
	channelName := c.Params.ByName("channelName")
	fmt.Printf("TODO: delete channel %s in room %s\n", channelName, name)

	c.String(200, "ok")
}
