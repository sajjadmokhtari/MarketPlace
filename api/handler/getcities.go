package handler

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"

	"github.com/gin-gonic/gin"
)

func GetCities(c *gin.Context) {
	var cities []model.City
	database := db.GetDb()
	if err := database.Find(&cities).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cities)
}
