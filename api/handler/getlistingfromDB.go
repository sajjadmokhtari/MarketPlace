package handler

import (
	"MarketPlace/cache"
	"MarketPlace/data/db"
	"MarketPlace/data/model"
	"MarketPlace/logging"
	"MarketPlace/pkg/metrics"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetListingsHandler نمایش همه آگهی‌ها بدون فیلتر
func GetListingsHandler(c *gin.Context) {
	listings, err := cache.GetListingsCache("ads:all")
	if err == nil {
		logging.GetLogger().Infow("✅ پاسخ از کش Redis ارسال شد", "key", "ads:all")
		c.JSON(http.StatusOK, listings)
		return
	} else {
		logging.GetLogger().Infow("ℹ️ کش Redis یافت نشد، خواندن از دیتابیس", "key", "ads:all")
	}

	var freshListings []model.Listing
	database := db.GetDb().Preload("City").Preload("Category")

	if err := database.Find(&freshListings).Error; err != nil {
		metrics.DbCall.WithLabelValues("find_listings", "fail").Inc()
		logging.GetLogger().Errorw("❌ خطا در خواندن آگهی‌ها از دیتابیس", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	metrics.DbCall.WithLabelValues("find_listings", "success").Inc()
	logging.GetLogger().Infow("✅ آگهی‌ها با موفقیت از دیتابیس خوانده شدند")

	// ذخیره در کش با اعتبار ۶۰ ثانیه
	if err := cache.SetListingsCache("ads:all", freshListings, time.Minute); err != nil {
		logging.GetLogger().Errorw("❌ خطا در ذخیره کش Redis", "error", err)
	} else {
		logging.GetLogger().Infow("📦 کش جدید ساخته شد و در Redis ذخیره شد", "key", "ads:all", "ttl", "60s")
	}

	// ارسال پاسخ به کاربر
	c.JSON(http.StatusOK, freshListings)
}
