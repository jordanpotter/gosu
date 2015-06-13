package accounts

import "github.com/gin-gonic/gin"

type CreateRequest struct {
	Email string `json:"email" form:"email" binding:"required"`
}

func (h *Handler) create(c *gin.Context) {
	var req CreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.AbortWithError(422, err)
		return
	}

	_, err = h.dbConn.CreateAccount(req.Email)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.String(200, "ok")
}
