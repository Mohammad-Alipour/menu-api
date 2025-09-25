package routes

import (
	"github.com/Mohammad-Alipour/menu-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	// Auth routes
	auth := api.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
	}

	// Public routes
	public := api.Group("/public")
	{
		public.GET("/categories", controllers.GetCategories)
		public.GET("/items", controllers.GetItems)
		public.GET("/items/:id", controllers.GetItemByID)
		public.GET("/items/:id/images", controllers.GetItemImages)
	}

	// Admin routes
	admin := api.Group("/admin")
	admin.Use(controllers.AdminAuthMiddleware())
	{
		// Categories
		admin.POST("/categories", controllers.CreateCategory)
		admin.PUT("/categories/:id", controllers.UpdateCategory)
		admin.DELETE("/categories/:id", controllers.DeleteCategory)

		// Items
		admin.POST("/items", controllers.CreateItem)
		admin.PUT("/items/:id", controllers.UpdateItem)
		admin.DELETE("/items/:id", controllers.DeleteItem)

		// Images
		admin.POST("/items/:id/images", controllers.UploadItemImage)
		admin.DELETE("/items/images/:id", controllers.DeleteItemImage)

		// Test route
		admin.GET("/secret", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You are admin 🚀"})
		})
	}
}
