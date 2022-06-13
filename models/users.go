package models

import (
	"time"

	"gorm.io/gorm"
)

type ContentUser struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey;auto_increment" json:"id"`
	Email          string    `gorm:"not null;unique" json:"email" `
	HashedPassword string    `gorm:"not null" json:"hashedpassword" `
	CreatedAt      time.Time // `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

//a login model displayed for  user to login
type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password" binding:"required,min=6"`
}

//returns a response when a User logs in
type UserResponse struct {
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

//this returns(generates) an acess token to the client
type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

//the new user response is used so that the password wont be exposed to the client
func NewUserResponse(response ContentUser) UserResponse {
	return UserResponse{
		Email:     response.Email,
		CreatedAt: response.CreatedAt,
	}
}
