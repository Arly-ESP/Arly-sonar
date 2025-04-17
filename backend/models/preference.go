package models

import "gorm.io/datatypes"

type Preferences struct {
	ID                 uint           `json:"id" gorm:"primary_key"`               
	UserID             uint           `json:"user_id" gorm:"not null"`              
	IsNotificationActive bool         `json:"is_notification_active" gorm:"default:true"` 
	ChatbotSettings    datatypes.JSON `json:"chatbot_settings" gorm:"type:json"`   
	User               User           `json:"user" gorm:"foreignkey:UserID"`      
}
