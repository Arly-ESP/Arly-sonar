package models 

import (
	"time"
)

type UserActivity struct {
	ID               uint         `json:"id" gorm:"primary_key"`
	UserID           uint         `json:"user_id" gorm:"not null"`       
	Date             time.Time    `json:"date" gorm:"not null"`         
	Mood 		  string       `json:"mood" gorm:"type:varchar(255);default:'neutral'"`
	MessageCount     int          `json:"message_count" gorm:"default:0"`
}

