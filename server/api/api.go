package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/api/v0"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/config/etcd"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/db/mongo"
)

var (
	port     int
	etcdAddr string
)

func init() {
	flag.IntVar(&port, "port", 8080, "the port to use")
	flag.StringVar(&etcdAddr, "etcd", "http://localhost:4001", "the etcd server addresses")
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	configConn := etcd.New([]string{etcdAddr})

	dbConn := getDBConn(configConn)
	defer dbConn.Close()

	tokenFactory := getTokenFactory(configConn)

	startServer(dbConn, tokenFactory)
}

func getDBConn(configConn config.Conn) *db.Conn {
	mongoConfig, err := configConn.GetMongo()
	if err != nil {
		panic(err)
	}

	dbConn, err := mongo.New(mongoConfig)
	if err != nil {
		panic(err)
	}
	return dbConn
}

func getTokenFactory(configConn config.Conn) *token.Factory {
	authTokenConfig, err := configConn.GetAuthToken()
	if err != nil {
		panic(err)
	}
	return token.NewFactory(authTokenConfig.Key, authTokenConfig.Duration)
}

func startServer(dbConn *db.Conn, tokenFactory *token.Factory) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v0Handler := v0.New(dbConn, tokenFactory)
	v0Handler.AddRoutes(r.Group("/v0"))

	r.Run(fmt.Sprintf(":%d", port))
}