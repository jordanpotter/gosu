package rooms

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type AuthenticationRequest struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type AuthenticationResponse struct {
	AuthToken   string    `json:"authToken"`
	AuthExpires time.Time `json:"authExpires"`
}

func (h *Handler) authenticate(c *gin.Context) {
	var req AuthenticationRequest
	if !c.Bind(&req) {
		return
	}

	accountID, err := c.Get(middleware.AccountIDKey)
	if err != nil {
		c.Fail(500, err)
		return
	}

	member, err := h.dbConn.Rooms.GetMemberByAccount(req.ID, accountID.(string))
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	fmt.Println("TODO: add room and admin info to auth token")
	fmt.Println(member)

	c.String(200, "ok")
}
