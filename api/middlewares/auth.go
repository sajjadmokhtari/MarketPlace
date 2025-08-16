package middlewares

import (
	"MarketPlace/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// گرفتن کوکی توکن
		token, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "توکن یافت نشد"})
			c.Abort()
			return
		}

		claims, err := services.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "توکن نامعتبر است"})
			c.Abort()
			return
		}

		// ذخیره شماره و نقش داخل کانتکست Gin
		c.Set("userPhone", claims.Phone)
		c.Set("userRole", claims.Role)

		c.Next()
	}
}





func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"message": "دسترسی فقط برای مدیران مجاز است"})
			c.Abort()
			return
		}
		c.Next()
	}
}
