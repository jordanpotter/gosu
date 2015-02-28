package main

import "fmt"
import "github.com/gin-gonic/gin"

import "github.com/JordanPotter/gosu-server/api/v0"

func main() {
	fmt.Println("api")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v0.AddRoutes(r.Group("/v0"))

	r.Run(":8080")
}
