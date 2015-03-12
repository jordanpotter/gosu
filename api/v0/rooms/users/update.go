package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) setAdmin(c *gin.Context) {
	var req SetAdminRequest
	if !c.Bind(&req) {
		return
	}

	roomName := c.Params.ByName("roomName")
	userName := c.Params.ByName("userName")
	fmt.Printf("TODO: set admin for user %s in room %s: %t\n", userName, roomName, req.Admin)

	c.String(200, "ok")
}

func (h *Handler) setBanned(c *gin.Context) {
	var req SetBannedRequest
	if !c.Bind(&req) {
		return
	}

	roomName := c.Params.ByName("roomName")
	userName := c.Params.ByName("userName")
	fmt.Printf("TODO: set banned for user %s in room %s: %t\n", userName, roomName, req.Banned)

	c.String(200, "ok")
}
