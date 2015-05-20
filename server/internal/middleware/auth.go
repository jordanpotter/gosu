package middleware

import (
	"errors"
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
			c.Fail(403, errors.New("token expired"))
			return
		}

		c.Set(TokenKey, t)
		c.Next()
	}
}
