package models

import (
	"time"
)

type UserAnswers struct {
	ID        uint                   `json:"id" gorm:"primaryKey"`
	UserID    uint                   `json:"user_id" gorm:"not null"`
	SurveyID  uint                   `json:"survey_id" gorm:"not null;index"`
	SurveySlug string                `json:"survey_slug" gorm:"not null;index"`	
	Answers   map[string]interface{} `json:"answers" gorm:"-"`         
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}