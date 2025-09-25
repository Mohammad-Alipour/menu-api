package controllers

import (
	"github.com/Mohammad-Alipour/menu-api/config"
	"github.com/Mohammad-Alipour/menu-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetItems Public
func GetItems(c *gin.Context) {
	var items []models.Item
	if err := config.DB.Preload("Category").Preload("Images").Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch items"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func GetItemByID(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.Preload("Category").Preload("Images").First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// CreateItem Admin
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create item"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	config.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Item{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
}
