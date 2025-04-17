
package models

type Context struct {
	ID                 uint                   `json:"id" gorm:"primary_key"`
	UserID             uint                   `json:"user_id" gorm:"not null"`
	Content            map[string]interface{} `json:"content" gorm:"type:json"`             
	IsCurrent          bool           `json:"is_current" gorm:"default:true"`
	ConversationHistory map[string]interface{} `json:"conversation_history" gorm:"type:json"`
	Intent             string                 `json:"intent"`
	Entities           map[string]interface{} `json:"entities" gorm:"type:json"`           
}
