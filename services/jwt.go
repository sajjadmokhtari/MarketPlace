package services

import (
	"crypto/rsa"
	"time"

	"MarketPlace/data/model"
	"MarketPlace/utils"

	"github.com/golang-jwt/jwt/v5"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

// بارگذاری کلیدها هنگام راه‌اندازی سرور
func InitJWTKeys(privatePath, publicPath string) error {
	var err error
	privateKey, err = utils.LoadPrivateKey(privatePath)
	if err != nil {
		return err
	}
	publicKey, err = utils.LoadPublicKey(publicPath)
	if err != nil {
		return err
	}
	return nil
}

// 🔹 این تابع یه توکن جدید می‌سازه برای کاربری که شماره تلفن و نقشش مشخصه
func GenerateJWT(phone, role string) (string, error) {
	claims := model.CustomClaims{
		Phone: phone,
		Role:  role,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        utils.GenerateJTI(),
			Issuer:    "your-app",
			Audience:  []string{"your-client"},
			Subject:   phone,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

// اعتبارسنجی JWT با کلید عمومی
func ValidateJWT(tokenStr string) (*model.CustomClaims, error) {
	claims := &model.CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
