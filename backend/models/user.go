package models

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primary_key auto_increment"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" omitEmpty:"true"`
	Email       string    `json:"email" validate:"required,email"`
	Password    string    `json:"password" validate:"required,min=8"`
	VerificationCode       string     `json:"verification_code,omitempty"`
	VerificationCodeExpiry *time.Time `json:"verification_code_expiry,omitempty"`
	Verified    bool      `json:"verified" gorm:"default:false"`
	IsDeleted   bool      `json:"is_deleted" gorm:"default:false"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at"`
	FirstSession bool     `json:"first_session" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
}
