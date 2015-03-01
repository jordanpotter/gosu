package main

import (
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

	port := viper.GetStringMapString("events")["port"]
	r.Run(":" + port)
}
