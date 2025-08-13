package router

import (
	"MarketPlace/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// سرو کردن فایل‌های استاتیک
	r.Static("/static", "./frontend")

	// API ها
	r.POST("/api/check-phone", handler.CheckPhoneHandler)
	r.POST("/api/send-otp", handler.SendOtpHandler)
	r.POST("/api/verify-otp", handler.VerifyOtpHandler)

	// داده‌ها
	r.GET("/api/categories", handler.GetCategories)
	r.GET("/api/cities", handler.GetCities)

	// صفحات HTML
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})
	r.GET("/dashboard", func(c *gin.Context) {
		c.File("./frontend/dashboard.html")
	})
	r.GET("/order", func(c *gin.Context) {
		c.File("./frontend/order.html")
	})

	return r
}
