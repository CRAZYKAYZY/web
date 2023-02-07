package main

import (
	"github.com/CRAZYKAYZY/web/cmd"
	"github.com/CRAZYKAYZY/web/db"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	//I want to be able to call this function and run the cobra command.
	cmd.Execute()

	db.ConnectDb()

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080")
}
