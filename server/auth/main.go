package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/auth/accounts"
	"github.com/jordanpotter/gosu/server/auth/rooms"
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
	defer configConn.Close()

	dbConn := getDBConn(configConn)
	defer dbConn.Close()

	tf := getTokenFactory(configConn)

	startServer(dbConn, tf)
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

func startServer(dbConn *db.Conn, tf *token.Factory) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	accountsHandler := accounts.New(dbConn, tf)
	accountsHandler.AddRoutes(r.Group("/accounts"))

	roomsHandler := rooms.New(dbConn, tf)
	roomsHandler.AddRoutes(r.Group("/rooms"))

	r.Run(fmt.Sprintf(":%d", port))
}
