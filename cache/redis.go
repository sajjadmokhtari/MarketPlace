package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // آدرس و پورت Redis
		Password: "",               // اگر پسورد داره بذارید
		DB:       0,
	})
}

func SetOTP(phone string, otp string) error {
	key := "otp:" + phone
	// 5 دقیقه انقضا
	return Client.Set(ctx, key, otp, 1*60*1000000000).Err()
}

func GetOTP(phone string) (string, error) {
	key := "otp:" + phone
	return Client.Get(ctx, key).Result()
}

func DeleteOTP(phone string) error {
	key := "otp:" + phone
	return Client.Del(ctx, key).Err()
}
