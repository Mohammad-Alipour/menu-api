package controllers

import (
	"github.com/Mohammad-Alipour/menu-api/config"
	"github.com/Mohammad-Alipour/menu-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetItemImages Public
func GetItemImages(c *gin.Context) {
	itemID := c.Param("id")
	var images []models.ItemImage
	if err := config.DB.Where("item_id = ?", itemID).Find(&images).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch item images"})
		return
	}
	c.JSON(http.StatusOK, images)
}

// UploadItemImage Admin
func UploadItemImage(c *gin.Context) {
	var image models.ItemImage
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := config.DB.Create(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not upload image"})
		return
	}
	c.JSON(http.StatusCreated, image)
}

func DeleteItemImage(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.ItemImage{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete image"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "image deleted"})
}
