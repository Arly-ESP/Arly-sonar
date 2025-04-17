package models

type UserPersonalityMBTI struct {
	ID              uint   `json:"id" gorm:"primary_key"`         
	UserID          uint   `json:"user_id" gorm:"not null"`      
	PersonalityType string `json:"personality_type" gorm:"not null"` 
	Description     string `json:"description" gorm:"type:text"` 
	User            User   `json:"user" gorm:"foreignkey:UserID"` 
}
