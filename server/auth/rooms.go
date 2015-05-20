package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type RoomAuthenticateRequest struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type RoomAuthenticateResponse struct {
	AuthToken   string    `json:"authToken"`
	AuthExpires time.Time `json:"authExpires"`
}

func createRoomAuthHandler(dbConn *db.Conn, tf *token.Factory) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RoomAuthenticateRequest
		if !c.Bind(&req) {
			return
		}

		accountID, err := c.Get(middleware.AccountIDKey)
		if err != nil {
			c.Fail(500, err)
			return
		}

		member, err := dbConn.Rooms.GetMemberByAccount(req.ID, accountID.(string))
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
}
