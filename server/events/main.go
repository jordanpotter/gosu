package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/config/etcd"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/db/mongo"
	"github.com/jordanpotter/gosu/server/internal/events"
	"github.com/jordanpotter/gosu/server/internal/events/nanomsg"
)

var (
	port     int
	etcdAddr string
)

func init() {
	flag.IntVar(&port, "port", 8082, "the port to use")
	flag.StringVar(&etcdAddr, "etcd", "http://localhost:4001", "the etcd server addresses")
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	configConn := etcd.New([]string{etcdAddr})
	defer configConn.Close()

	dbConn := getDBConn(configConn)
	defer dbConn.Close()

	sub := getSubscriber(configConn)
	defer sub.Close()

	go func() {
		recv := sub.Listen()
		for m := range recv {
			fmt.Println(m)
		}
	}()

	tf := getTokenFactory(configConn)

	startServer(dbConn, tf)
}

func getDBConn(configConn config.Conn) *db.Conn {
	mongoAddrs, err := configConn.GetMongoAddrs()
	if err != nil {
		panic(err)
	}

	mongoConfig, err := configConn.GetMongo()
	if err != nil {
		panic(err)
	}

	dbConn, err := mongo.New(mongoAddrs, mongoConfig)
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

func getSubscriber(configConn config.Conn) events.Subscriber {
	sub, err := nanomsg.NewSubscriber("127.0.0.1:9001")
	if err != nil {
		panic(err)
	}
	return sub
}

func startServer(dbConn *db.Conn, tf *token.Factory) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(fmt.Sprintf(":%d", port))
}
