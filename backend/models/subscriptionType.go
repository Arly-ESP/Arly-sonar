package models

type SubscriptionType struct {
	ID                uint    `json:"id" gorm:"primary_key"`        
	SubscriptionName  string  `json:"subscription_name" gorm:"not null"` 
	Price             float64 `json:"price" gorm:"not null"`        
	Duration          int     `json:"duration" gorm:"not null"`      
	IsRecurring       bool    `json:"is_recurring" gorm:"default:true"` 
	Subscriptions     []Subscription `json:"subscriptions" gorm:"foreignkey:SubscriptionTypeID"` 
}