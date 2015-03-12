package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) delete(c *gin.Context) {
	roomName := c.Params.ByName("roomName")
	userName := c.Params.ByName("userName")
	fmt.Printf("TODO: delete user %s in room %s\n", userName, roomName)

	c.String(200, "ok")
}
