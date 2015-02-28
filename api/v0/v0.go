package v0

import "github.com/gin-gonic/gin"

import "github.com/JordanPotter/gosu-server/api/v0/accounts"

func AddRoutes(rg *gin.RouterGroup) {
	accounts.AddRoutes(rg.Group("/accounts"))
}
