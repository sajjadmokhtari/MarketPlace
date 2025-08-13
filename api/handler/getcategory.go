package handler

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []model.Category
	database := db.GetDb()
	if err := database.Find(&categories).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, categories)
}
