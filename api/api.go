package main

import "fmt"
import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("api")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":8080")
}
