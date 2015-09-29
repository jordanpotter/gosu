package main

import (
	"flag"
	"fmt"

	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/config/etcd"
)

var (
	port     int
	etcdAddr string
)

func init() {
	flag.IntVar(&port, "port", 8083, "the port to use")
	flag.StringVar(&etcdAddr, "etcd", "http://localhost:4001", "the etcd server addresses")
	flag.Parse()
}

func main() {
	configConn := etcd.New([]string{etcdAddr})
	defer configConn.Close()

	tf := getTokenFactory(configConn)

	startServer(tf)
}

func getTokenFactory(configConn config.Conn) *token.Factory {
	authTokenConfig, err := configConn.GetAuthToken()
	if err != nil {
		panic(err)
	}
	return token.NewFactory(authTokenConfig.Key, authTokenConfig.Duration)
}

func startServer(tf *token.Factory) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(fmt.Sprintf(":%d", port))
}
