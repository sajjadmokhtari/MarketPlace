package middlewares

import (
	"fmt"
	"strconv"
	"time"

	"MarketPlace/pkg/metrics" // مسیر واقعی metrics رو بذار

	"github.com/gin-gonic/gin"
)

// PrometheusMiddleware همه متریک‌های HTTP، DB و Login رو ثبت می‌کنه
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()                     // ⏱ زمان شروع درخواست
		c.Next()                                // ▶️ اجرای Handler اصلی و ادامه میدل‌ویر
		duration := time.Since(start).Seconds() // ⏱ محاسبه زمان پاسخ

		path := c.FullPath() // مسیر ثبت شده در Router
		if path == "" {
			path = c.Request.URL.Path // مسیر واقعی URL اگر FullPath خالی بود
		}

		method := c.Request.Method  // متد HTTP (GET/POST/...)
		status := c.Writer.Status() // کد وضعیت HTTP

		// 📝 چاپ لاگ برای دیباگ
		fmt.Printf("📊 Prometheus Log: path=%q method=%s status=%d duration=%.2fms\n", path, method, status, duration)

		// 📊 ثبت زمان پاسخ HTTP در Prometheus
		metrics.HttpDuration.WithLabelValues(path, method, strconv.Itoa(status)).
			Observe(duration) // ثبت زمان پاسخ

		// 🔢 ثبت تعداد درخواست‌ها (کل/نا موفق)
		var opStatus string
		if status >= 200 && status < 400 {
			opStatus = "success" // موفق
		} else {
			opStatus = "fail" // ناموفق
		}
		metrics.DbCall.WithLabelValues("http_request", opStatus).Inc() // افزایش Counter

		// 🔑 مثال لاگین: ثبت تعداد تلاش‌های موفق/ناموفق
		if path == "/login" {
			statusVal, exists := c.Get("login_status") // گرفتن وضعیت Login
			if exists {
				statusStr := fmt.Sprintf("%v", statusVal)
				metrics.LoginAttempts.WithLabelValues(statusStr).Inc() // افزایش Counter
			}
		}
	}
}
