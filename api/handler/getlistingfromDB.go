package handler

import (
	"MarketPlace/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetListingsHandler فعلاً فقط نمایش همه آگهی‌ها بدون فیلتر دسته‌بندی
func GetListingsHandler(c *gin.Context) {
	listings, err := services.GetAllListings("") // دسته‌بندی خالی یعنی همه آگهی‌ها
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, listings)
}
