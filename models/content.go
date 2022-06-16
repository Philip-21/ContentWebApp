package models

import (
	"time"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	ID       uint   `gorm:"primarykey"`
	Title    string `gorm:"not null" json:"title"`
	Contents string `gorm:"not null" json:"contents"`
	Comment  string `gorm:"not null" json:"comment"`
	// //OwnerID is the fkey that refers  the ContentUserID
	// OwnerID ContentUser `gorm:"foreignKey:ContentRefer"`
	// //ContentRefer is the column that defines  the foreignkey
	// ContentRefer uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
