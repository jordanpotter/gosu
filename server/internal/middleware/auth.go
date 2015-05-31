package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
)

const (
	TokenKey   = "token"
	authHeader = "Authorization"
)

func AuthRequired(tf *token.Factory) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get(authHeader)
		t, err := tf.Decrypt(auth)
		if err != nil {
			c.Fail(500, err)
			return
		} else if t.Expires.Before(time.Now()) {
			c.Fail(401, errors.New("token expired"))
			return
		}

		c.Set(TokenKey, t)
		c.Next()
	}
}

func AuthMatchesRoom(roomIDParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t, err := c.Get(TokenKey)
		if err != nil {
			c.Fail(500, err)
			return
		}

		roomID := c.Params.ByName(roomIDParam)
		if roomID == "" {
			c.Fail(403, errors.New("invalid room id"))
			return
		}

		authRoomID := t.(*token.Token).Room.ID
		if roomID != authRoomID {
			c.Fail(403, fmt.Errorf("room id %s does not match auth token's room %s", roomID, authRoomID))
			return
		}

		c.Next()
	}
}

func IsRoomAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, err := c.Get(TokenKey)
		if err != nil {
			c.Fail(500, err)
			return
		}

		admin := t.(*token.Token).Room.Admin
		if !admin {
			c.Fail(403, errors.New("must be admin for room"))
			return
		}

		c.Next()
	}
}

func IsNotSameMember(memberIDParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t, err := c.Get(TokenKey)
		if err != nil {
			c.Fail(500, err)
			return
		}

		memberID := c.Params.ByName(memberIDParam)
		if memberID == "" {
			c.Fail(403, errors.New("invalid member id"))
			return
		}

		authMemberID := t.(*token.Token).Room.MemberID
		if memberID == authMemberID {
			c.Fail(403, fmt.Errorf("member id %s cannot match auth token's member id %s", memberID, authMemberID))
			return
		}

		c.Next()
	}
}
