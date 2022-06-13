package models

import (
	"time"

	"gorm.io/gorm"
)

//this creates a table in the database
func MigrateContent(db *gorm.DB) error {
	err := db.AutoMigrate(&Content{}, &ContentUser{})
	return err
}

type Content struct {
	gorm.Model
	ID        uint   `gorm:"primarykey"`
	Title     string `gorm:"not null" json:"title"`
	Contents  string `gorm:"not null" json:"contents"`
	Comment   string `gorm:"null" json:"comment"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
