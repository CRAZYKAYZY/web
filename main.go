package main

import (
	"github.com/CRAZYKAYZY/web/db"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	db.ConnectDb()

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080")
}
