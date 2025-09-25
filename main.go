package main

import (
	"github.com/Mohammad-Alipour/menu-api/config"
	"github.com/Mohammad-Alipour/menu-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		panic("❌ Server failed to start: " + err.Error())
	}
}
