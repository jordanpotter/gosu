package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

const port = 1337

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	startServer()
}

func startServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(":1337")
}
