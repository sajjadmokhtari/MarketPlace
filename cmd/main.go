package main

import (
	"MarketPlace/api/router"
	"MarketPlace/cache"
	"MarketPlace/data/db"
	"MarketPlace/data/db/migration"
	"log"
)

func main() {
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
	log.Println("🚀 سرور روی پورت 8080 اجرا شد")
	log.Fatal(r.Run(":8080"))
}
