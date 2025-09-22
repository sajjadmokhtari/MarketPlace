package handler

import (
	"MarketPlace/logging"
	"MarketPlace/pkg/metrics"
	"MarketPlace/services"
	"MarketPlace/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyOtpHandler(c *gin.Context) {
    var req OTPRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        logging.GetLogger().Errorw("Error decoding OTP request", "error", err)
        c.JSON(http.StatusBadRequest, Response{Valid: false, Message: "درخواست نامعتبر"})
        metrics.LoginAttempts.WithLabelValues("fail").Inc()
        return
    }

    if err := services.VerifyOTP(req.Phone, req.OTP); err != nil {
        logging.GetLogger().Errorw("OTP verification failed", "error", err, "phone", req.Phone)
        c.JSON(http.StatusBadRequest, Response{Valid: false, Message: err.Error()})
        metrics.LoginAttempts.WithLabelValues("fail").Inc()
        return
    }

    role := services.UserRole(req.Phone)

    if _, err := services.UpdateOrCreateUser(req.Phone, role); err != nil {
        logging.GetLogger().Errorw("Error creating or updating user", "error", err, "phone", req.Phone)
        c.JSON(http.StatusInternalServerError, Response{Valid: false, Message: "خطا در ثبت کاربر"})
        metrics.LoginAttempts.WithLabelValues("fail").Inc()
        return
    }

    // 📌 گرفتن userID از دیتابیس
    user, err := services.GetUserByPhone(req.Phone)
    if err != nil {
        logging.GetLogger().Errorw("Error fetching user after creation", "error", err, "phone", req.Phone)
        c.JSON(http.StatusInternalServerError, Response{Valid: false, Message: "خطا در دریافت اطلاعات کاربر"})
        metrics.LoginAttempts.WithLabelValues("fail").Inc()
        return
    }

    // ساخت Access Token با userID
    accessToken, err := services.GenerateJWT(user.ID, req.Phone, role)
    if err != nil {
        logging.GetLogger().Errorw("Error generating access token", "error", err, "phone", req.Phone)
        c.JSON(http.StatusInternalServerError, Response{Valid: false, Message: "خطا در ساخت توکن دسترسی"})
        metrics.LoginAttempts.WithLabelValues("fail").Inc()
        return
    }

    // ساخت Refresh Token
    refreshToken, err := services.GenerateRefreshToken(req.Phone)
    if err != nil {
        logging.GetLogger().Errorw("Error generating refresh token", "error", err, "phone", req.Phone)
        c.JSON(http.StatusInternalServerError, Response{Valid: false, Message: "خطا در ساخت رفرش توکن"})
        metrics.LoginAttempts.WithLabelValues("fail").Inc()
        return
    }

    // ذخیره توکن‌ها در کوکی امن
    utils.SetAuthCookie(c, accessToken, 900)        // Access Token: 15 دقیقه
    utils.SetRefreshCookie(c, refreshToken, 604800) // Refresh Token: 7 روز

    logging.GetLogger().Infow("User authenticated and tokens issued", "phone", req.Phone)
    metrics.LoginAttempts.WithLabelValues("success").Inc()

    c.JSON(http.StatusOK, Response{Valid: true, Message: "ورود موفق، توکن‌ها صادر شدند"})
}
