package members

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/events/types"
	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
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

	roomID := c.Params.ByName("roomID")
	room, err := h.dbConn.Rooms.Get(roomID)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	passwordMatches := password.MatchesHash(req.Password, room.PasswordHash)
	if !passwordMatches {
		c.AbortWithError(403, errors.New("invalid password"))
		return
	}

	accountID := authToken.Account.ID
	member, err := h.dbConn.Rooms.AddMember(roomID, accountID, req.Name)
	if err == db.DuplicateError {
		c.AbortWithError(409, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	e := &types.RoomMemberCreated{
		RoomID:     roomID,
		MemberID:   member.ID,
		MemberName: member.Name,
		Admin:      member.Admin,
		Banned:     member.Banned,
		Created:    member.Created,
	}
	err = h.pub.Send(e)
	if err != nil {
		fmt.Printf("Failed to send event: %v", err)
	}

	c.JSON(200, member)
}
