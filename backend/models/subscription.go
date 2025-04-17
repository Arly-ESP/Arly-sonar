package models

import "time"

type Subscription struct {
	ID                 uint      `json:"id" gorm:"primary_key"`        
	DateStarted        time.Time `json:"date_started" gorm:"not null"` 
	DateEnded          *time.Time `json:"date_ended,omitempty"`        
	UserID             uint      `json:"user_id" gorm:"not null"`      
	SubscriptionTypeID uint      `json:"subscription_type_id" gorm:"not null"`
	User               User      `json:"user" gorm:"foreignkey:UserID"` 
	SubscriptionType   SubscriptionType `json:"subscription_type" gorm:"foreignkey:SubscriptionTypeID"`
}
