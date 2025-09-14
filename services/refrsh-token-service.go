package services

import (
	"errors"
	"time"

	"MarketPlace/cache"
	"MarketPlace/utils"
)

// ساخت Refresh Token و ذخیره در Redis
func GenerateRefreshToken(phone string) (string, error) {
    token := utils.GenerateJTI() // استفاده از تابع مشترک برای UUID
    key := "refresh:" + token    // حالا توکن می‌شه کلید، و phone می‌شه مقدار

    err := cache.Client.Set(ctx, key, phone, 7*24*time.Hour).Err()
    if err != nil {
        return "", err
    }
    return token, nil
}

// اعتبارسنجی Refresh Token از Redis
func ValidateRefreshToken(token string) (string, error) {
    key := "refresh:" + token
    phone, err := cache.Client.Get(ctx, key).Result()
    if err != nil {
        return "", errors.New("رفرش توکن نامعتبر است")
    }
    return phone, nil
}
