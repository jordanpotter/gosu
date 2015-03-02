package main

import (
	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu-server/api/v0"

	"github.com/JordanPotter/gosu-server/internal/db"
	"github.com/JordanPotter/gosu-server/internal/db/mongo"
)

const address = ":8080"

func main() {
	dbConn, err := mongo.New()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	startServer(dbConn)
}

func startServer(dbConn db.Conn) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v0Handler := v0.New(dbConn)
	v0Handler.AddRoutes(r.Group("/v0"))

	r.Run(address)
}
