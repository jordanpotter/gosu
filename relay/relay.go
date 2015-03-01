package main

import (
	"github.com/gin-gonic/gin"
)

const address = ":8082"

func main() {
	startServer()
}

func startServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(address)
}
