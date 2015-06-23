package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/config/etcd"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/db/postgres"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub/nanomsg"
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
	configConn := etcd.New([]string{etcdAddr})
	defer configConn.Close()

	dbConn := getDBConn(configConn)
	defer dbConn.Close()

	sub := getSubscriber(configConn)
	defer sub.Close()

	tf := getTokenFactory(configConn)

	startServer(dbConn, tf, sub)
}

func getDBConn(configConn config.Conn) db.Conn {
	postgresAddrs, err := configConn.GetPostgresAddrs()
	if err != nil {
		panic(err)
	}

	postgresConfig, err := configConn.GetPostgres()
	if err != nil {
		panic(err)
	}

	dbConn, err := postgres.New(postgresAddrs, postgresConfig)
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

func getSubscriber(configConn config.Conn) pubsub.Subscriber {
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

func startServer(dbConn db.Conn, tf *token.Factory, sub pubsub.Subscriber) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// hub, err := hub.New(tf, sub)
	// if err != nil {
	// 	panic(err)
	// }
	// hub.AddRoutes(r.Group("/events"))

	r.Run(fmt.Sprintf(":%d", port))
}
