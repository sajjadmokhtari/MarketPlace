package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // اینو امن نگه دار، نذار توی کد پابلیک باشه

type CustomClaims struct {
	Phone string `json:"phone"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(phone string, role string) (string, error) {
	claims := CustomClaims{
		Phone: phone,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // اعتبار ۲۴ ساعته
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// اعتبارسنجی توکن و استخراج اطلاعات
func ValidateJWT(tokenStr string) (*CustomClaims, error) {
    claims := &CustomClaims{}

    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("توکن نامعتبر است")
    }

    return claims, nil
}

