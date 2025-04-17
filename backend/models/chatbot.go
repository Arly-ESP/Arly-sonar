package models

type Chatbot struct {
	ID      uint    `json:"id" gorm:"primary_key"`
	Name    string  `json:"name" gorm:"not null"`
	Version string  `json:"version" gorm:"not null"`
	Contexts []Context `json:"contexts" gorm:"foreignkey:ChatbotID"` 
}
