package main

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type AccountAuthenticateRequest struct {
	Email          string `json:"email" form:"email" binding:"required"`
	DeviceName     string `json:"deviceName" form:"deviceName" binding:"required"`
	DevicePassword string `json:"devicePassword" form:"devicePassword" binding:"required"`
}

type AccountAuthenticateResponse struct {
	AuthToken   string    `json:"authToken"`
	AuthExpires time.Time `json:"authExpires"`
}

func createAccountAuthHandler(dbConn *db.Conn, tf *token.Factory) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AccountAuthenticateRequest
		if !c.Bind(&req) {
			return
		}

		account, err := dbConn.Accounts.GetByEmail(req.Email)
		if err == db.NotFoundError {
			c.Fail(403, err)
			return
		} else if err != nil {
			c.Fail(500, err)
			return
		}

		if !hasValidDeviceCredentials(req.DeviceName, req.DevicePassword, account.Devices) {
			c.Fail(403, errors.New("no matching device name and password"))
			return
		}

		authToken := tf.New()
		authToken.Account.ID = account.ID
		authTokenEncrypted, err := tf.Encrypt(authToken)
		if err != nil {
			c.Fail(500, err)
			return
		}

		c.JSON(200, AccountAuthenticateResponse{authTokenEncrypted, authToken.Expires})
	}
}

func hasValidDeviceCredentials(deviceName, devicePassword string, devices []db.Device) bool {
	for _, device := range devices {
		if deviceName == device.Name && password.MatchesHash(devicePassword, device.PasswordHash) {
			return true
		}
	}
	return false
}
