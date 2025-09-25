package main

import (
	"github.com/Mohammad-Alipour/menu-api/config"
	"github.com/Mohammad-Alipour/menu-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// Run seeder once to insert initial data
	Seed()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
