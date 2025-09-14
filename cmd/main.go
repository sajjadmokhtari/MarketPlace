package main

import (
	"MarketPlace/api/router"
	"MarketPlace/cache"
	"MarketPlace/data/db"
	"MarketPlace/data/db/migration"
	"MarketPlace/logging" // اضافه شد: پکیج لاگر خودمون
	"MarketPlace/pkg/metrics"
	"MarketPlace/services"
)

func main() {
	// 🚀 راه‌اندازی Logger
	logging.InitLogger()
	log := logging.GetLogger()

	// اتصال به DB
	if err := db.InitDb(); err != nil {
		log.Fatalf("❌ failed to connect to database: %v", err)
	}

	// مهاجرت و داده‌های پیش‌فرض
	migration.Up_1()

	// راه‌اندازی Redis
	cache.InitRedis()

	// بار گذاری کلید ها برای  استفاده در امضا  توکن
	if err := services.InitJWTKeys("keys/private.pem", "keys/public.pem"); err != nil {
		log.Fatalf("❌ خطا در بارگذاری کلیدهای JWT: %v", err)
	}

	// ثبت متریک‌ها
	metrics.RegisterAll()

	// اتصال به MongoDB
	if err := db.InitMongo(); err != nil {
		log.Fatalf("❌ failed to connect to MongoDB: %v", err)
	}

	// ثبت مسیرها
	r := router.SetupRoutes()

	log.Infow("🚀 سرور روی پورت 8080 اجرا شد")

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("❌ سرور نتوانست اجرا شود: %v", err)
	}
}
