package main

import (
	"github.com/JordanPotter/gosu-server/api/v0"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.ReadInConfig()

	startServer()
}

func startServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v0.AddRoutes(r.Group("/v0"))

	port := viper.GetStringMapString("api")["port"]
	r.Run(":" + port)
}
