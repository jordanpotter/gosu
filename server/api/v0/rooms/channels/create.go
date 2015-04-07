package channels

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func (h *Handler) create(c *gin.Context) {
	var req CreateRequest
	if !c.Bind(&req) {
		return
	}

	roomName := c.Params.ByName("roomName")
	fmt.Printf("TODO: create channel %s in %s if doesn't exist\n", req.Name, roomName)

	c.String(200, "ok")
}
