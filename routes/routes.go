package routes

import (
	"github.com/Mohammad-Alipour/menu-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
		}
		admin := api.Group("/admin")
		{
			admin.GET("/secret", controllers.AdminAuthMiddleware(), func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "You are admin 🚀"})
			})
		}
	}
}
