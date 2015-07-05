package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/api/v0"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/config/etcd"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/db/postgres"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub/nanomsg"
)

var (
	port           int
	pubPort        int
	etcdAddr       string
	migrationsPath string
)

func init() {
	flag.IntVar(&port, "port", 8081, "the port to use")
	flag.IntVar(&pubPort, "publisher port", 9001, "the port to publish events on")
	flag.StringVar(&etcdAddr, "etcd", "http://localhost:4001", "the etcd server addresses")
	flag.StringVar(&migrationsPath, "migrations", "conf/db/migrations", "the path to migration assets")
	flag.Parse()
}

func main() {
	configConn := etcd.New([]string{etcdAddr})
	defer configConn.Close()

	dbConn := getDBConn(configConn)
	defer dbConn.Close()

	pub := getPublisher(configConn)
	defer pub.Close()

	tf := getTokenFactory(configConn)

	startServer(dbConn, tf, pub)
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

	dbConn, err := postgres.New(postgresAddrs, postgresConfig, migrationsPath)
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

func getPublisher(configConn config.Conn) pubsub.Publisher {
	pub, err := nanomsg.NewPublisher(fmt.Sprintf(":%d", pubPort))
	if err != nil {
		panic(err)
	}
	return pub
}

func startServer(dbConn db.Conn, tf *token.Factory, pub pubsub.Publisher) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v0Handler := v0.New(dbConn, tf, pub)
	v0Handler.AddRoutes(r.Group("/v0"))

	r.Run(fmt.Sprintf(":%d", port))
}
