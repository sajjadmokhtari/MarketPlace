package handler

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
	"MarketPlace/pkg/metrics"

	"github.com/gin-gonic/gin"
)

func GetCities(c *gin.Context) {
	var cities []model.City
	database := db.GetDb()
	if err := database.Find(&cities).Error; err != nil {
		metrics.DbCall.WithLabelValues("get_Cities", "fail").Inc()
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	metrics.DbCall.WithLabelValues("get_Cities", "success").Inc()
	c.JSON(200, cities)
}
