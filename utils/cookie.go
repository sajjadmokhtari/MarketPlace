package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetAuthCookie(c *gin.Context, token string, maxAge int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   false, // چون روی localhost هستی، باید false باشه
		SameSite: http.SameSiteStrictMode,
	})
}

func GetAuthCookie(c *gin.Context) (string, error) {
	return c.Cookie("auth_token")
}

func SetRefreshCookie(c *gin.Context, token string, maxAge int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Path:     "/api/refresh", // محدود به مسیر رفرش
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   false, // روی سرور واقعی باید true بشه
		SameSite: http.SameSiteStrictMode,
	})
}

func GetRefreshCookie(c *gin.Context) (string, error) {
	return c.Cookie("refresh_token")
}
