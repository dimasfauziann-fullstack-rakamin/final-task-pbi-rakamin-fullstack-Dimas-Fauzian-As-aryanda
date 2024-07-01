package helpers

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(userId uint) (string, error) {
    expirationTime := time.Now().Add(72 * time.Hour)
    claims := &jwt.StandardClaims{
        ExpiresAt: expirationTime.Unix(),
        Id:        string(userId),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    return tokenString, err
}

func ValidateJWT(tokenString string) (uint, error) {
    claims := &jwt.StandardClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return 0, err
    }
    if !token.Valid {
        return 0, err
    }
    return uint(claims.Id), nil
}
