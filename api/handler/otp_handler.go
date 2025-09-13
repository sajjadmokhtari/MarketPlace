package handler

import (
	"MarketPlace/logging"
	"MarketPlace/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
