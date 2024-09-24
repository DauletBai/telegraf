package models

import "time"

type Message struct {
    ID         int       `json:"id"`
    SenderID   int       `json:"sender_id"`
    ReceiverID int       `json:"receiver_id"`
    Content    string    `json:"content"`
    MessageType string   `json:"message_type"` // Текст, аудио, видео
    Timestamp  time.Time `json:"timestamp"`
}