package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/config"
)

var (
	port        int
	etcdAddress string
)

func init() {
	flag.IntVar(&port, "port", 8081, "the port to use")
	flag.StringVar(&etcdAddress, "etcd", "http://127.0.0.1:4001", "the etcd server address")
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
