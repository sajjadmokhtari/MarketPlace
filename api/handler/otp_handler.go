package handler

import (
	"MarketPlace/logging"
	"MarketPlace/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendOtpHandler هندلر ارسال OTP

// SendOtpHandler هندلر ارسال OTP
func SendOtpHandler(c *gin.Context) {
	var req PhoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.GetLogger().Errorw("Error decoding phone request", "error", err)
		c.JSON(http.StatusBadRequest, Response{Valid: false, Message: "درخواست نامعتبر"})
		return
	}

	logging.GetLogger().Infow("SendOtpHandler received phone", "phone", req.Phone)

	if err := services.SendOTP(req.Phone); err != nil {
		logging.GetLogger().Errorw("Error sending OTP", "error", err, "phone", req.Phone)
		c.JSON(http.StatusInternalServerError, Response{Valid: false, Message: "خطا در ارسال OTP"})
		return
	}

	c.JSON(http.StatusOK, Response{Valid: true, Message: "کد OTP ارسال شد"})
}

// VerifyOtpHandler هندلر بررسی OTP
func VerifyOtpHandler(c *gin.Context) {
	var req OTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.GetLogger().Errorw("Error decoding OTP request", "error", err)
		c.JSON(http.StatusBadRequest, Response{Valid: false, Message: "درخواست نامعتبر"})
		return
	}

	if err := services.VerifyOTP(req.Phone, req.OTP); err != nil {
		logging.GetLogger().Errorw("OTP verification failed", "error", err, "phone", req.Phone)
		c.JSON(http.StatusBadRequest, Response{Valid: false, Message: err.Error()})
		return
	}

	// ساخت JWT بعد از تایید OTP
	token, err := services.GenerateJWT(req.Phone, "user") // نقش فعلاً "user"
	if err != nil {
		logging.GetLogger().Errorw("Error generating JWT", "error", err, "phone", req.Phone)
		c.JSON(http.StatusInternalServerError, Response{Valid: false, Message: "خطا در ساخت توکن"})
		return
	}

	logging.GetLogger().Infow("Generated JWT for phone", "phone", req.Phone, "token", token)

	// ذخیره توکن در کوکی
	c.SetCookie("token", token, 3600, "/", "", false, true)

	// پاسخ موفقیت
	c.JSON(http.StatusOK, Response{Valid: true, Message: "شما با موفقیت وارد شدید، توکن ذخیره شد"})
}
