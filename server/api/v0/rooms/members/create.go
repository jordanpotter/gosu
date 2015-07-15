package members

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/api/v0/sanitization"
	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/events"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type JoinRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (h *Handler) join(c *gin.Context) {
	var req JoinRequest
	err := c.Bind(&req)
	if err != nil {
		c.AbortWithError(422, err)
		return
	}

	t, ok := c.Get(middleware.TokenKey)
	if !ok {
		c.AbortWithError(500, errors.New("missing auth token"))
		return
	}
	authToken := t.(*token.Token)

	roomIDString := c.Params.ByName("roomID")
	roomID, err := strconv.Atoi(roomIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	room, err := h.dbConn.GetRoom(roomID)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	passwordMatches := password.MatchesBcryptHash(req.Password, room.PasswordHash)
	if !passwordMatches {
		c.AbortWithError(403, errors.New("invalid password"))
		return
	}

	member, err := h.dbConn.CreateMember(authToken.Account.ID, roomID, req.Name)
	if err == db.DuplicateError {
		c.AbortWithError(409, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = h.pub.Send(events.RoomMemberCreated{
		RoomID:     roomID,
		MemberID:   member.ID,
		MemberName: member.Name,
		Admin:      member.Admin,
		Banned:     member.Banned,
		Created:    member.Created,
	})
	if err != nil {
		fmt.Printf("Failed to send event: %v", err)
	}

	c.JSON(200, sanitization.ToMember(member))
}
