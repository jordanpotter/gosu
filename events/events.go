package main

import (
	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu-server/internal/config"
)

func main() {
	config, err := config.Get()
	if err != nil {
		panic(err)
	}

	startServer(&config.Events)
}

func startServer(eventsConfig *config.Events) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(eventsConfig.Address)
}
