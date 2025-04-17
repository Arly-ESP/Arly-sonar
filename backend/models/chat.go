package models

type Chat struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	UserID    uint    `json:"user_id" gorm:"not null"`
	ContextID uint    `json:"context_id" gorm:"not null"`
	Context   Context   `json:"context"`
}
