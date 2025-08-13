package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") 

type CustomClaims struct {
	Phone string `json:"phone"`
	Role  string `json:"role"`
	jwt.RegisteredClaims //  شامل فیلد های مهم مثل زمان انقضا و زمان 
}

func GenerateJWT(phone string, role string) (string, error) {
	claims := CustomClaims{
		Phone: phone,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // توکن تا ۲۴ ساعت اعتبار داره
			IssuedAt:  jwt.NewNumericDate(time.Now()),//زمان صدور توکن.
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // ساخت توکن با الگوریتم  اچ اس 
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// اعتبارسنجی توکن و استخراج اطلاعات
func ValidateJWT(tokenStr string) (*CustomClaims, error) {
    claims := &CustomClaims{}//آماده‌سازی ساختار برای استخراج اطلاعات

    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("توکن نامعتبر است")
    }

    return claims, nil
}

