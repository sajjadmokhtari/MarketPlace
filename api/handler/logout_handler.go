package handler

import (
	"net/http"
	"time"

	// فرض بر اینه که لاگرت اینجاست
	"MarketPlace/logging"
	"MarketPlace/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogoutHandler(c *gin.Context) {
	tokenStr, err := c.Cookie("auth_token")
	if err != nil {
		logging.GetLogger().Infow("توکن در کوکی یافت نشد", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "توکن یافت نشد"})
		return
	}

	claims, err := services.ValidateJWT(tokenStr)
	if err != nil {

		logging.GetLogger().Infow("توکن نامعتبر در logout", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "توکن نامعتبر"})
		return
	}

	jti := claims.ID
	exp := claims.ExpiresAt.Time
	ttl := time.Until(exp)

	logging.GetLogger().Infow("در حال بلاک کردن توکن", zap.String("jti", jti), zap.Duration("ttl", ttl))

	err = services.BlacklistToken(jti, ttl)
	if err != nil {

		logging.GetLogger().Infow("خطا در بلاک کردن توکن", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بلاک کردن توکن"})
		return
	}

	// حذف کوکی‌ها
	c.SetCookie("auth_token", "", -1, "/", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)

	logging.GetLogger().Infow("کاربر با موفقیت logout شد", zap.String("phone", claims.Phone), zap.String("jti", jti))

	c.JSON(http.StatusOK, gin.H{"message": "✅ خروج با موفقیت انجام شد"})
}
