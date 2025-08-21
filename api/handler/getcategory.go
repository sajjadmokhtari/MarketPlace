package handler

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
	"MarketPlace/pkg/metrics"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []model.Category
	database := db.GetDb()
	if err := database.Find(&categories).Error; err != nil {
		metrics.DbCall.WithLabelValues("get_categories", "fail").Inc()

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	metrics.DbCall.WithLabelValues("get_categories", "success").Inc()

	c.JSON(200, categories)
}
