package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu/server/api/v0"
	"github.com/JordanPotter/gosu/server/internal/auth/token"
	"github.com/JordanPotter/gosu/server/internal/config"
	"github.com/JordanPotter/gosu/server/internal/db"
	"github.com/JordanPotter/gosu/server/internal/db/mongo"
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

	dbConn, err := mongo.New(&config.DB.Mongo)
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	tokenFactory := token.NewFactory(config.Auth.Token.Key, config.Auth.Token.Duration)

	startServer(dbConn, tokenFactory, &config.API)
}

func startServer(dbConn *db.Conn, tokenFactory *token.Factory, apiConfig *config.API) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v0Handler := v0.New(dbConn, tokenFactory)
	v0Handler.AddRoutes(r.Group("/v0"))

	r.Run(fmt.Sprintf(":%d", apiConfig.Port))
}
