package utils

import (
    "crypto/rsa"
    "errors"
    "os"

    "github.com/golang-jwt/jwt/v5"
)

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
    keyData, err := os.ReadFile(path)
    if err != nil {
        return nil, errors.New("خطا در خواندن کلید خصوصی: " + err.Error())
    }
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
    if err != nil {
        return nil, errors.New("خطا در پارس کلید خصوصی: " + err.Error())
    }
    return privateKey, nil
}

func LoadPublicKey(path string) (*rsa.PublicKey, error) {
    keyData, err := os.ReadFile(path)
    if err != nil {
        return nil, errors.New("خطا در خواندن کلید عمومی: " + err.Error())
    }
    publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
    if err != nil {
        return nil, errors.New("خطا در پارس کلید عمومی: " + err.Error())
    }
    return publicKey, nil
}
