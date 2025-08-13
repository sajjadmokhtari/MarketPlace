package handler

import (
	"MarketPlace/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendOtpHandler هندلر ارسال OTP
func SendOtpHandler(c *gin.Context) {
	var req PhoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error decoding phone request:", err)
		c.JSON(http.StatusBadRequest, Response{Valid: false, Message: "درخواست نامعتبر"})
		return
	}

	log.Printf("SendOtpHandler received phone: %s", req.Phone)

	if err := services.SendOTP(req.Phone); err != nil {
		log.Println("Error sending OTP:", err)
		c.JSON(http.StatusInternalServerError, Response{Valid: false, Message: "خطا در ارسال OTP"})
		return
	}

	c.JSON(http.StatusOK, Response{Valid: true, Message: "کد OTP ارسال شد"})
}

// VerifyOtpHandler هندلر بررسی OTP
func VerifyOtpHandler(c *gin.Context) {
	var req OTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error decoding OTP request:", err)
		c.JSON(http.StatusBadRequest, Response{Valid: false, Message: "درخواست نامعتبر"})
		return
	}

	if err := services.VerifyOTP(req.Phone, req.OTP); err != nil {
		log.Println("OTP verification failed:", err)
		c.JSON(http.StatusBadRequest, Response{Valid: false, Message: err.Error()})
		return
	}

	// ساخت JWT بعد از تایید OTP
	token, err := services.GenerateJWT(req.Phone, "user") // نقش فعلاً "user"
	if err != nil {
		log.Println("Error generating JWT:", err)
		c.JSON(http.StatusInternalServerError, Response{Valid: false, Message: "خطا در ساخت توکن"})
		return
	}

	log.Printf("Generated JWT for phone %s: %s", req.Phone, token)

	// ذخیره توکن در کوکی
	c.SetCookie("token", token, 3600, "/", "", false, true)

	// پاسخ موفقیت
	c.JSON(http.StatusOK, Response{Valid: true, Message: "شما با موفقیت وارد شدید، توکن ذخیره شد"})
}
