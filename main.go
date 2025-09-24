package main

import (
	"github.com/Mohammad-Alipour/menu-api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	//ConnectDatabase
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
