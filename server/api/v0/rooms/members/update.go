package members

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/api/middleware"
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

func (h *Handler) authenticate(c *gin.Context) {
	accountID, err := c.Get(middleware.AccountIDKey)
	if err != nil {
		c.Fail(500, err)
		return
	}

	roomName := c.Params.ByName("roomName")
	member, err := h.dbConn.Rooms.GetMember(roomName, accountID.(string))
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	fmt.Println("TODO: add room and admin info to auth token")
	fmt.Println("TODO: add account to channel")
	fmt.Println(member)

	c.String(200, "ok")
}

func (h *Handler) setAdmin(c *gin.Context) {
	var req SetAdminRequest
	if !c.Bind(&req) {
		return
	}

	fmt.Println("TODO: check account is admin for room")
	fmt.Println("TODO: check not revoking admin for self")

	roomName := c.Params.ByName("roomName")
	accountID := c.Params.ByName("memberAccountID")
	err := h.dbConn.Rooms.SetMemberAdmin(roomName, accountID, req.Admin)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
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
	accountID := c.Params.ByName("memberAccountID")
	err := h.dbConn.Rooms.SetMemberBanned(roomName, accountID, req.Banned)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
