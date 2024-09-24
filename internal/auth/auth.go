package auth

import (
    "fmt"
    "telegraf/internal/models"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtSecret = []byte("секретный_ключ")

// Функция для генерации JWT токена для авторизованного пользователя
func GenerateToken(user *models.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })
    return token.SignedString(jwtSecret)
}

// Функция для проверки JWT токена
func ParseToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtSecret, nil
    })
    return token, err
}