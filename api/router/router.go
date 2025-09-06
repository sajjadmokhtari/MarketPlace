package router

import (
	"MarketPlace/api/handler"
	"MarketPlace/api/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.PrometheusMiddleware())

	// سرو کردن فایل‌های استاتیک (CSS, JS, تصاویر)
	r.Static("/static", "./frontend")
	r.Static("/uploads", "./uploads")

	// API ها
	api := r.Group("/api")
	{
		api.POST("/check-phone", handler.CheckPhoneHandler)
		api.POST("/send-otp", handler.SendOtpHandler)
		api.POST("/verify-otp", handler.VerifyOtpHandler)

		// ثبت و گرفتن آگهی‌ها
		api.POST("/listings", middlewares.AuthMiddleware(), handler.CreateListingHandler)
		api.GET("/listings", handler.GetListingsHandler)

		api.GET("/search", handler.SearchListingsHandler)

		// داده‌ها
		api.GET("/categories", handler.GetCategories)
		api.GET("/cities", handler.GetCities)
	}

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
	r.GET("/listings", func(c *gin.Context) {
		c.File("./frontend/listings.html")
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))// وقتی پرومتوس بیاد تمام متریک های ثبت شده رو نشون میده

	return r
}
