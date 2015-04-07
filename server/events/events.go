package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu/server/internal/config"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "conf/server.yaml", "Specify the configuration file path")
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	config, err := config.Load(configPath)
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

	r.Run(fmt.Sprintf(":%d", eventsConfig.Port))
}
