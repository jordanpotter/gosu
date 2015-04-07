package rooms

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Password string `json:"password" form:"password" binding:"required"`
}

type SetPasswordRequest struct {
	Password string `json:"password" form:"password" binding:"required"`
}

func (h *Handler) login(c *gin.Context) {
	var req LoginRequest
	if !c.Bind(&req) {
		return
	}

	roomName := c.Params.ByName("roomName")
	fmt.Printf("TODO: login to room %s\n", roomName)

	c.String(200, "ok")
}

func (h *Handler) logout(c *gin.Context) {
	roomName := c.Params.ByName("roomName")
	fmt.Printf("TODO: logout of room %s\n", roomName)

	c.String(200, "ok")
}

func (h *Handler) setPassword(c *gin.Context) {
	var req SetPasswordRequest
	if !c.Bind(&req) {
		return
	}

	roomName := c.Params.ByName("roomName")
	fmt.Printf("TODO: set password for room %s\n", roomName)

	c.String(200, "ok")
}
