package main

import (
	"MarketPlace/api/router"
	"MarketPlace/cache"
	"MarketPlace/data/db"
	"MarketPlace/data/db/migration"
	"MarketPlace/logging" // اضافه شد: پکیج لاگر خودمون
)

func main() {
	// 🚀 راه‌اندازی Logger
	logging.InitLogger()       // فقط یک بار لازمه
	log := logging.GetLogger() // گرفتن logger برای استفاده راحت

	// اتصال به DB
	if err := db.InitDb(); err != nil {
		log.Fatalf("❌ failed to connect to database: %v", err)
	}

	// مهاجرت و داده‌های پیش‌فرض
	migration.Up_1()

	// راه‌اندازی Redis
	cache.InitRedis()

	// ثبت مسیرها و گرفتن Engine
	r := router.SetupRoutes()

	// اجرای سرور Gin
	log.Infow("🚀 سرور روی پورت 8080 اجرا شد")

	// اگر سرور نتوانست اجرا شود
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ سرور نتوانست اجرا شود: %v", err)
	}
}


