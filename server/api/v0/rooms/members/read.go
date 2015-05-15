package members

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAll(c *gin.Context) {
	roomID := c.Params.ByName("roomID")
	fmt.Printf("TODO: get users in room %s\n", roomID)
	c.String(200, "ok")
}
