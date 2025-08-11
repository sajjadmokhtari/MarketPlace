package router

import (
	"MarketPlace/api/handler"
	"net/http"
)

func SetupRoutes() {
    // سرو کردن فایل‌های استاتیک روی مسیر /static/
    fs := http.FileServer(http.Dir("./frontend"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // API ها
    http.HandleFunc("/api/check-phone", handler.CheckPhoneHandler)
    http.HandleFunc("/api/send-otp", handler.SendOtpHandler)
    http.HandleFunc("/api/verify-otp", handler.VerifyOtpHandler)

    // هندل کردن صفحه اصلی
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./frontend/index.html")
    })
}

