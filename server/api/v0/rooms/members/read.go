package members

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAll(c *gin.Context) {
	roomName := c.Params.ByName("roomName")
	fmt.Printf("TODO: get users in room %s\n", roomName)
	c.String(200, "ok")
}
