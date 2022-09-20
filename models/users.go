package models

import (
	"time"

	"gorm.io/gorm"
)

type ContentUser struct {
	gorm.Model
	ID        int       `gorm:"primaryKey"`
	Email     string    `gorm:"not null;unique" json:"email" `
	FirstName string    `gorm:"not null" json:"firstname" `
	LastName  string    `gorm:"not null" json:"lastname" `
	Password  []byte    `gorm:"size:60;not null" json:"password" `
	CreatedAt time.Time // `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type SignInResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
