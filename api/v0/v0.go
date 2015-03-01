package v0

import (
	"github.com/JordanPotter/gosu-server/api/v0/accounts"
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	accounts.AddRoutes(rg.Group("/accounts"))
}
