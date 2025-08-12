package main

import (
	"MarketPlace/api/router"
	"MarketPlace/cache"
	"log"
	"net/http"
)

func main() {

	// راه‌اندازی Redis و Router
	cache.InitRedis()
	router.SetupRoutes()

	log.Println("🚀 سرور روی پورت 8080 اجرا شد")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// اضافه کن بالا
