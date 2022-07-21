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
	////Owner is the fkey that refers  the ContentUserID
	//Owner ContentUser `gorm:"foreignKey:ContentRefer"`
	////ContentRefer is the column that defines  the foreignkey
	//ContentRefer string `json:"UsersEmail"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
