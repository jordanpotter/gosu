package members

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

	fmt.Println("TODO: check account is admin for room")
	fmt.Println("TODO: check not revoking admin for self")

	roomName := c.Params.ByName("roomName")
	memberName := c.Params.ByName("memberName")
	err := h.dbConn.Rooms.SetMemberAdmin(roomName, memberName, req.Admin)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}

func (h *Handler) setBanned(c *gin.Context) {
	var req SetBannedRequest
	if !c.Bind(&req) {
		return
	}

	fmt.Println("TODO: check account is admin for room")
	fmt.Println("TODO: check not trying to ban self")

	roomName := c.Params.ByName("roomName")
	memberName := c.Params.ByName("memberName")
	err := h.dbConn.Rooms.SetMemberBanned(roomName, memberName, req.Banned)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
