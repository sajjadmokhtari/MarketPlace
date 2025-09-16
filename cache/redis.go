package cache

import (
	"context"
	"strconv"
	"time"

	"MarketPlace/logging"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // حتما IPv4 بذار، با localhost بعضی وقتا IPv6 می‌ره
		Password: "pass",           // اگه پسورد نداری خالی بذار ""
		DB:       0,
	})

	// تست اتصال
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		logging.GetLogger().Fatalw("❌ Failed to connect to Redis", "error", err)
	} else {
		logging.GetLogger().Infow("✅ Connected to Redis successfully", "addr", "127.0.0.1:6379")
	}
}

// ذخیره OTP با انقضای ۱ دقیقه
func SetOTP(phone string, otp string) error {
	key := "otp:" + phone
	err := Client.Set(ctx, key, otp, 2*time.Minute).Err()
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to save OTP in Redis", "phone", phone, "error", err)
	}
	return err
}

func GetOTP(phone string) (string, error) {
	key := "otp:" + phone
	val, err := Client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		logging.GetLogger().Errorw("❌ Failed to get OTP from Redis", "phone", phone, "error", err)
	}
	return val, err
}

func DeleteOTP(phone string) error {
	key := "otp:" + phone
	err := Client.Del(ctx, key).Err()
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to delete OTP from Redis", "phone", phone, "error", err)
	}
	return err
}

// بررسی اینکه آیا مجاز به ارسال OTP هست یا نه (حداقل ۶۰ ثانیه فاصله)
func CanSendOTP(phone string) bool {
	key := "otp:last:" + phone
	lastTimeStr, err := Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return true
	}
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to check CanSendOTP", "phone", phone, "error", err)
		return true
	}
	lastTime, err := strconv.ParseInt(lastTimeStr, 10, 64)
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to parse last OTP timestamp", "phone", phone, "error", err)
		return true
	}
	return time.Now().Unix()-lastTime >= 30 // ← تغییر فاصله به ۳۰ ثانیه
}

func MarkOTPSent(phone string) {
	key := "otp:last:" + phone
	err := Client.Set(ctx, key, time.Now().Unix(), 1*time.Hour).Err() // ← TTL برای ثبت زمان آخرین ارسال
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to mark OTP sent", "phone", phone, "error", err)
	}
}

// شمارش تعداد درخواست‌های OTP در ۱۰ دقیقه اخیر
func OTPRequestCount(phone string) int {
	key := "otp:count:" + phone
	countStr, err := Client.Get(ctx, key).Result()//مقدار شمارش رو  میخونه
	if err == redis.Nil {
		return 0
	}
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to get OTP request count", "phone", phone, "error", err)
		return 0
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to parse OTP request count", "phone", phone, "error", err)
		return 0
	}
	return count
}

func IncrementOTPRequest(phone string) {
	key := "otp:count:" + phone
	err := Client.Incr(ctx, key).Err()
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to increment OTP request count", "phone", phone, "error", err)
	}
	Client.Expire(ctx, key, 1*time.Hour) // ← تغییر TTL به ۱ ساعت
}

func IncrementFailedAttempts(phone string) {
	key := "otp:fail:" + phone
	err := Client.Incr(ctx, key).Err()
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to increment failed attempts", "phone", phone, "error", err)
	}
	Client.Expire(ctx, key, 1*time.Hour) // ← TTL برای شمارش تلاش‌ها
}

func GetFailedAttempts(phone string) int {
	key := "otp:fail:" + phone
	countStr, err := Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return 0
	}
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to get failed attempts", "phone", phone, "error", err)
		return 0
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to parse failed attempts", "phone", phone, "error", err)
		return 0
	}
	return count
}

func BlockPhone(phone string) {
	key := "otp:block:" + phone
	err := Client.Set(ctx, key, "true", 15*time.Minute).Err() // ← بلاک موقت ۱۵ دقیقه
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to block phone", "phone", phone, "error", err)
	}
}

func IsPhoneBlocked(phone string) bool {
	key := "otp:block:" + phone
	_, err := Client.Get(ctx, key).Result()
	return err != redis.Nil
}

func ResetFailedAttempts(phone string) {
	key := "otp:fail:" + phone
	err := Client.Del(ctx, key).Err()
	if err != nil {
		logging.GetLogger().Errorw("❌ Failed to reset failed attempts", "phone", phone, "error", err)
	}
}




