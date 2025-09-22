package services

import (
    "context"
    "time"

    "MarketPlace/cache"
    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// ذخیره jti در لیست سیاه با TTL برابر با زمان انقضای توکن
func BlacklistToken(jti string, ttl time.Duration) error {
    key := "blacklist:" + jti
    return cache.Client.Set(ctx, key, "true", ttl).Err()
}

// بررسی اینکه آیا jti در لیست سیاه هست
func IsTokenBlacklisted(jti string) (bool, error) {
    key := "blacklist:" + jti
    val, err := cache.Client.Get(ctx, key).Result()

    if err == redis.Nil {
        return false, nil // یعنی در لیست سیاه نیست
    }
    if err != nil {
        return false, err // خطای ارتباط با Redis
    }
    return val == "true", nil
}
