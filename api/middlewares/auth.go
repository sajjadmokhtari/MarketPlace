package middlewares

import (
	"MarketPlace/services"
	"context"
	"net/http"
)

type contextKey string

const (
	UserPhoneKey contextKey = "userPhone"
	UserRoleKey  contextKey = "userRole"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "توکن یافت نشد", http.StatusUnauthorized)
			return
		}

		claims, err := services.ValidateJWT(cookie.Value)
		if err != nil {
			http.Error(w, "توکن نامعتبر است", http.StatusUnauthorized)
			return
		}

		// ذخیره اطلاعات در context
		ctx := context.WithValue(r.Context(), UserPhoneKey, claims.Phone)
		ctx = context.WithValue(ctx, UserRoleKey, claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}



func AdminMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        role, ok := r.Context().Value(UserRoleKey).(string)
        if !ok || role != "admin" {
            http.Error(w, "دسترسی فقط برای مدیران مجاز است", http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}
