package services

import (
	"errors"
	"time"

	"MarketPlace/cache"

	"github.com/google/uuid"
)

// ساخت Refresh Token و ذخیره در Redis
func GenerateRefreshToken(phone string) (string, error) {
	token := uuid.NewString()
	key := "refresh:" + phone
	err := cache.Client.Set(ctx, key, token, 7*24*time.Hour).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

// اعتبارسنجی Refresh Token از Redis
func ValidateRefreshToken(token string) (string, error) {
	keys := cache.Client.Keys(ctx, "refresh:*").Val()
	for _, key := range keys {
		val := cache.Client.Get(ctx, key).Val()
		if val == token {
			phone := key[len("refresh:"):]
			return phone, nil
		}
	}
	return "", errors.New("رفرش توکن نامعتبر است")
}
