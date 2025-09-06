package handler

import (
	"MarketPlace/data/db"
	"MarketPlace/data/model"
	"MarketPlace/pkg/metrics"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetListingsHandler نمایش همه آگهی‌ها بدون فیلتر
func GetListingsHandler(c *gin.Context) {
	var listings []model.Listing

	database := db.GetDb().Preload("City").Preload("Category")

	// کوئری مستقیم به دیتابیس
	if err := database.Find(&listings).Error; err != nil {
		metrics.DbCall.WithLabelValues("find_listings", "fail").Inc()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metrics.DbCall.WithLabelValues("find_listings", "success").Inc()
	c.JSON(http.StatusOK, listings)
}
