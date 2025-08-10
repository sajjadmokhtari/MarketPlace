package router

import (
	"MarketPlace/api/handler"
	"net/http"
)

func SetupRoutes() {
	// سرو کردن کل فولدر frontend به صورت استاتیک
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	// API ها
	http.HandleFunc("/api/check-phone", handler.CheckPhoneHandler)
	http.HandleFunc("/api/send-otp", handler.SendOtpHandler)
	http.HandleFunc("/api/verify-otp", handler.VerifyOtpHandler)
}
