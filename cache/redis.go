package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "pass",
		DB:       0,
	})
}

// ذخیره OTP با انقضای ۵ دقیقه
func SetOTP(phone string, otp string) error {
	key := "otp:" + phone
	return Client.Set(ctx, key, otp, 5*time.Minute).Err()
}

func GetOTP(phone string) (string, error) {
	key := "otp:" + phone
	return Client.Get(ctx, key).Result()
}

func DeleteOTP(phone string) error {
	key := "otp:" + phone
	return Client.Del(ctx, key).Err()
}

// بررسی اینکه آیا مجاز به ارسال OTP هست یا نه (حداقل ۶۰ ثانیه فاصله)
func CanSendOTP(phone string) bool {
	key := "otp:last:" + phone
	lastTimeStr, err := Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return true
	}
	lastTime, err := strconv.ParseInt(lastTimeStr, 10, 64)
	if err != nil {
		return true
	}
	return time.Now().Unix()-lastTime >= 60
}

// ثبت زمان ارسال OTP
func MarkOTPSent(phone string) {
	key := "otp:last:" + phone
	Client.Set(ctx, key, time.Now().Unix(), 10*time.Minute)
}

// شمارش تعداد درخواست‌های OTP در ۱۰ دقیقه اخیر
func OTPRequestCount(phone string) int {
	key := "otp:count:" + phone
	countStr, err := Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return 0
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0
	}
	return count
}

// افزایش شمارش درخواست OTP
func IncrementOTPRequest(phone string) {
	key := "otp:count:" + phone
	Client.Incr(ctx, key)
	Client.Expire(ctx, key, 10*time.Minute)
}
