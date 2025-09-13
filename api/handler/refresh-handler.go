package handler

import (
    "MarketPlace/logging"
    "MarketPlace/services"
    "MarketPlace/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)

func RefreshTokenHandler(c *gin.Context) {
    refreshToken, err := utils.GetRefreshCookie(c)
    if err != nil || refreshToken == "" {
        logging.GetLogger().Warnw("Missing refresh token in cookie")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "رفرش توکن یافت نشد"})
        return
    }

    phone, err := services.ValidateRefreshToken(refreshToken)
    if err != nil {
        logging.GetLogger().Warnw("Invalid refresh token", "error", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "رفرش توکن نامعتبر است"})
        return
    }

    newAccessToken, err := services.GenerateJWT(phone, "user")
    if err != nil {
        logging.GetLogger().Errorw("Failed to generate new access token", "error", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ساخت توکن جدید"})
        return
    }

    utils.SetAuthCookie(c, newAccessToken, 900) // 15 دقیقه

    logging.GetLogger().Infow("Access token refreshed", "phone", phone)
    c.JSON(http.StatusOK, gin.H{"message": "توکن جدید صادر شد"})
}
