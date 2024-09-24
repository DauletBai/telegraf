package auth

import (
    "net/http"
    "strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Токен не предоставлен", http.StatusUnauthorized)
            return
        }

        tokenString := strings.Split(authHeader, "Bearer ")[1]
        token, err := ParseToken(tokenString)
        if err != nil || !token.Valid {
            http.Error(w, "Недействительный токен", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}