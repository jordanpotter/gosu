package accounts

import "errors"

import "github.com/gin-gonic/gin"

import "github.com/JordanPotter/gosu-server/api/v0/accounts/auth"

type CreateRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CreateResponse struct {
	Id   string    `json:"id"`
	Auth auth.Auth `json:"auth"`
}

func create(c *gin.Context) {
	var req CreateRequest
	if !c.Bind(&req) {
		return
	}

	// TODO: create account in database
	id := "id"

	auth, err := auth.New(id)
	if err != nil {
		c.Fail(500, errors.New("accounts: failed to create authentication data"))
		return
	}

	c.JSON(200, CreateResponse{id, auth})
}
