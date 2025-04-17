package models

import (
	"time"
)

type Message struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Date         time.Time `json:"date" gorm:"not null"`
	MessageType  string    `json:"message_type"` 
	IsBotMessage bool      `json:"is_bot_message"`
	ResponseTime int       `json:"response_time"`
	Content      string    `json:"content" gorm:"type:text;not null"` 
	ChatID       uint      `json:"chat_id" gorm:"not null"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	Chat         Chat      `json:"chat"` 
	User         User      `json:"user" gorm:"foreignkey:UserID"`
}
