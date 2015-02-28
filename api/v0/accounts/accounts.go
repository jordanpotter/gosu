package accounts

import "github.com/gin-gonic/gin"

func AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", create)
}
