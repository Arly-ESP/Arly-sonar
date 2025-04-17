package models

import "time"

type TrainingData struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	TrainingContent string    `json:"training_content" gorm:"type:text;not null"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ChatbotID       uint      `json:"chatbot_id"` 
	Chatbot         Chatbot   `json:"chatbot" gorm:"foreignkey:ChatbotID"` 
}
