package rooms

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) get(c *gin.Context) {
	roomName := c.Params.ByName("roomName")
	fmt.Printf("TODO: get room %s\n", roomName)

	c.String(200, "ok")
}
