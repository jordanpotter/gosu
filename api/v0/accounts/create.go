package accounts

import "github.com/gin-gonic/gin"

type CreateRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CreateResponse struct {
	Id          string `json:"id"`
	EventsToken string `json:"eventsToken"`
}

func create(c *gin.Context) {
	var req CreateRequest
	if !c.Bind(&req) {
		return
	}

	// TODO: create account in database

	res := CreateResponse{"id", "token"}
	c.JSON(200, res)
}
