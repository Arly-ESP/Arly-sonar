package models

import (
	"time"
)

type Question struct {
	Question        string                 `json:"question"`
	QuestionType    string                 `json:"question_type"`
	QuestionOptions map[string]interface{} `json:"question_options,omitempty"`
	Order           int                    `json:"order"`
}

type Surveys struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	SurveyName        string    `json:"survey_name" gorm:"not null"`
	SurveySlug 	 string    	`json:"survey_slug" gorm:"not null;unique"`
	SurveyDescription string    `json:"survey_description" gorm:"not null"`
	Questions         string    `json:"questions" gorm:"type:json"` 
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}