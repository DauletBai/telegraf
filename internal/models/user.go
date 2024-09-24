package models

import "time"

type User struct {
    ID           int       `json:"id"`
    Name         string    `json:"name"`
    Phone        string    `json:"phone"`
    PasswordHash string    `json:"-"`  // Не отправляется в JSON ответах
    CreatedAt    time.Time `json:"created_at"`
}