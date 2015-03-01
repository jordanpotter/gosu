package main

import (
	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu-server/api/v0"
)

const address = ":8080"

func main() {
	startServer()
}

func startServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v0.AddRoutes(r.Group("/v0"))

	r.Run(address)
}
