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

type SetAdminRequest struct {
	// Cannot use `binding:"required"` here, since the validation
	// check will fail when `admin` is false.
	Admin bool `json:"admin" form:"admin"`
}

type SetBannedRequest struct {
	// Cannot use `binding:"required"` here, since the validation
	// check will fail when `banned` is false.
	Banned bool `json:"banned" form:"banned"`
}

func (h *Handler) login(c *gin.Context) {
	var req LoginRequest
	if !c.Bind(&req) {
		return
	}

	name := c.Params.ByName("name")
	fmt.Printf("TODO: login to room %s\n", name)

	c.String(200, "ok")
}

func (h *Handler) logout(c *gin.Context) {
	name := c.Params.ByName("name")
	fmt.Printf("TODO: logout of room %s\n", name)

	c.String(200, "ok")
}

func (h *Handler) setPassword(c *gin.Context) {
	var req SetPasswordRequest
	if !c.Bind(&req) {
		return
	}

	name := c.Params.ByName("name")
	fmt.Printf("TODO: set password for room %s\n", name)

	c.String(200, "ok")
}

func (h *Handler) setAdmin(c *gin.Context) {
	var req SetAdminRequest
	if !c.Bind(&req) {
		return
	}

	name := c.Params.ByName("name")
	userName := c.Params.ByName("userName")
	fmt.Printf("TODO: set admin for user %s in room %s: %t\n", userName, name, req.Admin)

	c.String(200, "ok")
}

func (h *Handler) setBanned(c *gin.Context) {
	var req SetBannedRequest
	if !c.Bind(&req) {
		return
	}

	name := c.Params.ByName("name")
	userName := c.Params.ByName("userName")
	fmt.Printf("TODO: set banned for user %s in room %s: %t\n", userName, name, req.Banned)

	c.String(200, "ok")
}

func (h *Handler) moveChannel(c *gin.Context) {
	name := c.Params.ByName("name")
	channelName := c.Params.ByName("channelName")
	fmt.Printf("TODO: move to channel %s for room %s\n", channelName, name)

	c.String(200, "ok")
}
