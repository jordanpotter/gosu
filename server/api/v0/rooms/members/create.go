package members

import (
	"errors"

	"github.com/gin-gonic/gin"
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
	if !c.Bind(&req) {
		return
	}

	t, err := c.Get(middleware.TokenKey)
	if err != nil {
		c.Fail(500, err)
		return
	}

	roomID := c.Params.ByName("roomID")
	room, err := h.dbConn.Rooms.Get(roomID)
	if err != nil {
		c.Fail(500, err)
		return
	}

	passwordMatches := password.MatchesHash(req.Password, room.PasswordHash)
	if !passwordMatches {
		c.Fail(403, errors.New("invalid password"))
		return
	}

	accountID := t.(*token.Token).Account.ID
	err = h.dbConn.Rooms.AddMember(roomID, accountID, req.Name)
	if err == db.DuplicateError {
		c.Fail(409, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
