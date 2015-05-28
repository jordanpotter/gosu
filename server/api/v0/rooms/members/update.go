package members

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/events"
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

	fmt.Println("TODO: check not revoking admin for self")

	roomID := c.Params.ByName("roomID")
	memberID := c.Params.ByName("memberID")
	err := h.dbConn.Rooms.SetMemberAdmin(roomID, memberID, req.Admin)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	e := events.RoomMemberAdminUpdated{
		RoomID:   roomID,
		MemberID: memberID,
		Admin:    req.Admin,
	}
	err = h.pub.Send(e)
	if err != nil {
		fmt.Println("Failed to send event: %v", err)
	}

	c.String(200, "ok")
}

func (h *Handler) setBanned(c *gin.Context) {
	var req SetBannedRequest
	if !c.Bind(&req) {
		return
	}

	fmt.Println("TODO: check not trying to ban self")

	roomID := c.Params.ByName("roomID")
	memberID := c.Params.ByName("memberID")
	err := h.dbConn.Rooms.SetMemberBanned(roomID, memberID, req.Banned)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	e := events.RoomMemberBannedUpdated{
		RoomID:   roomID,
		MemberID: memberID,
		Banned:   req.Banned,
	}
	err = h.pub.Send(e)
	if err != nil {
		fmt.Println("Failed to send event: %v", err)
	}

	c.String(200, "ok")
}
