package models

type Gamification struct {
	ID         uint `json:"id" gorm:"primary_key"`    
	UserID     uint `json:"user_id" gorm:"not null"`  
	StreakCount int  `json:"streak_count" gorm:"default:0"` 
	User       User `json:"user" gorm:"foreignkey:UserID"` 
}
