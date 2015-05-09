package rooms

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/api/middleware"
	"github.com/jordanpotter/gosu/server/internal/auth/password"
)

type LoginRequest struct {
	Password string `json:"password" form:"password" binding:"required"`
	PeerName string `json:"peerName" form:"peerName" binding:"required"`
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
	room, err := h.dbConn.Rooms.GetByName(roomName)
	if err != nil {
		c.Fail(500, err)
		return
	}

	fmt.Println("TODO: check if account has membership already. If so, do not check password")

	passwordMatches := password.MatchesHash(req.Password, room.PasswordHash)
	if !passwordMatches {
		c.Fail(403, errors.New("invalid password"))
		return
	}

	accountId, err := c.Get(middleware.AccountIdKey)
	if err != nil {
		c.Fail(500, err)
		return
	}

	err = h.dbConn.Accounts.AddMembership(accountId.(string), room.Id, req.PeerName)
	if err != nil {
		c.Fail(500, err)
		return
	}

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
