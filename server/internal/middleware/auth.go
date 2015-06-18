package middleware

import (
	"errors"
	"fmt"
	"strconv"
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
			c.AbortWithError(500, err)
			return
		} else if t.Expires.Before(time.Now()) {
			c.AbortWithError(401, errors.New("token expired"))
			return
		}
		c.Set(TokenKey, t)
		c.Next()
	}
}

func AuthMatchesRoom(roomIDParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t, ok := c.Get(TokenKey)
		if !ok {
			c.AbortWithError(403, errors.New("missing auth token"))
			return
		}

		roomIDString := c.Params.ByName(roomIDParam)
		roomID, err := strconv.Atoi(roomIDString)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		authRoomID := t.(*token.Token).Room.ID
		if roomID != authRoomID {
			c.AbortWithError(403, fmt.Errorf("room id %d does not match auth token's room %d", roomID, authRoomID))
			return
		}
		c.Next()
	}
}

func IsRoomAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, ok := c.Get(TokenKey)
		if !ok {
			c.AbortWithError(403, errors.New("missing auth token"))
			return
		}

		admin := t.(*token.Token).Room.Admin
		if !admin {
			c.AbortWithError(403, errors.New("must be admin for room"))
			return
		}
		c.Next()
	}
}

func IsNotSameMember(memberIDParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t, ok := c.Get(TokenKey)
		if !ok {
			c.AbortWithError(403, errors.New("missing auth token"))
			return
		}

		memberIDString := c.Params.ByName(memberIDParam)
		memberID, err := strconv.Atoi(memberIDString)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		authMemberID := t.(*token.Token).Room.MemberID
		if memberID == authMemberID {
			c.AbortWithError(403, fmt.Errorf("member id %d cannot match auth token's member id %d", memberID, authMemberID))
			return
		}
		c.Next()
	}
}
