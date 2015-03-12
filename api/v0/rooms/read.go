package rooms

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) get(c *gin.Context) {
	name := c.Params.ByName("name")
	fmt.Printf("TODO: get room %s\n", name)

	c.String(200, "ok")
}

func (h *Handler) getUsers(c *gin.Context) {
	name := c.Params.ByName("name")
	fmt.Printf("TODO: get users in room %s\n", name)

	c.String(200, "ok")
}

func (h *Handler) getRelayConns(c *gin.Context) {
	name := c.Params.ByName("name")
	channelName := c.Params.ByName("channelName")
	fmt.Printf("TODO: get relay connections for channel %s in room %s\n", channelName, name)

	c.String(200, "ok")
}
