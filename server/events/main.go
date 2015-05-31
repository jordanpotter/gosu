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
	subPort  int
	etcdAddr string
)

func init() {
	flag.IntVar(&port, "port", 8082, "the port to use")
	flag.IntVar(&subPort, "subscriber port", 9002, "the port to subscribe to events from")
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

	listenChan := make(chan *events.Message, 100)
	err := sub.Listen(listenChan)
	if err != nil {
		panic(err)
	}

	go func() {
		for m := range listenChan {
			fmt.Println(m.Event)
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
	apiAddrs, err := configConn.GetAPIAddrs()
	if err != nil {
		panic(err)
	}

	sub, err := nanomsg.NewSubscriber()
	if err != nil {
		panic(err)
	}

	subAddrs := make([]string, 0, len(apiAddrs))
	for _, apiAddr := range apiAddrs {
		addr := fmt.Sprintf("%s:%d", apiAddr.IP.String(), apiAddr.PubPort)
		subAddrs = append(subAddrs, addr)
	}
	err = sub.SetAddrs(subAddrs)
	if err != nil {
		sub.Close()
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
