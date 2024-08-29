package auth

import (
    "github.com/dgrijalva/jwt-go"
    "milo-ia/internal/config"
    "time"
)

var secretKey string

func InitJWT(cfg config.Config) {
    secretKey = cfg.JWTSecret
}

func GenerateToken(username string) (string, error) {
    claims := &jwt.StandardClaims{
        ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        Issuer:    username,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenStr string) (*jwt.StandardClaims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })
    if err != nil {
        return nil, err
    }
    claims, ok := token.Claims.(*jwt.StandardClaims)
    if !ok || !token.Valid {
        return nil, err
    }
    return claims, nil
}
