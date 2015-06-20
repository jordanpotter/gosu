package members

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/api/v0/sanitization"
	"github.com/jordanpotter/gosu/server/internal/db"
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
	err := c.Bind(&req)
	if err != nil {
		c.AbortWithError(422, err)
		return
	}

	roomIDString := c.Params.ByName("roomID")
	roomID, err := strconv.Atoi(roomIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	memberIDString := c.Params.ByName("memberID")
	memberID, err := strconv.Atoi(memberIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	member, err := h.dbConn.SetMemberAdminForRoom(memberID, roomID, req.Admin)
	if err == db.NotFoundError {
		c.AbortWithError(404, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// e := &types.RoomMemberAdminUpdated{
	// 	RoomID:   roomID,
	// 	MemberID: memberID,
	// 	Admin:    req.Admin,
	// }
	// err = h.pub.Send(e)
	// if err != nil {
	// 	fmt.Printf("Failed to send event: %v", err)
	// }

	c.JSON(200, sanitization.ToMember(member))
}

func (h *Handler) setBanned(c *gin.Context) {
	var req SetBannedRequest
	err := c.Bind(&req)
	if err != nil {
		c.AbortWithError(422, err)
		return
	}

	roomIDString := c.Params.ByName("roomID")
	roomID, err := strconv.Atoi(roomIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	memberIDString := c.Params.ByName("memberID")
	memberID, err := strconv.Atoi(memberIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	member, err := h.dbConn.SetMemberBannedForRoom(memberID, roomID, req.Banned)
	if err == db.NotFoundError {
		c.AbortWithError(404, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// e := &types.RoomMemberBannedUpdated{
	// 	RoomID:   roomID,
	// 	MemberID: memberID,
	// 	Banned:   req.Banned,
	// }
	// err = h.pub.Send(e)
	// if err != nil {
	// 	fmt.Printf("Failed to send event: %v", err)
	// }

	c.JSON(200, sanitization.ToMember(member))
}
