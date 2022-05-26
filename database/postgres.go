package database

import (
	"errors"

	"github.com/Philip-21/proj1/models"
	"gorm.io/gorm"
)

///var db *gorm.DB

func GetContents(db *gorm.DB) ([]models.Content, error) {
	contents := []models.Content{}
	query := db.Select("contents.*").Group("content_id")
	err := query.Find(&contents).Error
	if err != nil {
		return contents, err
	}
	return contents, nil
}

func GetContentByID(id string, db *gorm.DB) (models.Content, bool, error) {
	c := models.Content{}

	query := db.Select("books.*")
	query = query.Group("contents.id")
	err := query.Where("contents.id = ?", id).First(&c).Error
	err1 := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !err1 {
		return c, false, err
	}
	if err1 {
		return c, false, nil
	}
	return c, true, nil
}

func DeleteContent(id string, db *gorm.DB) error {
	var b models.Content
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateContent(db *gorm.DB, b *models.Content) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}

//////////Users
//func CreateUser (db*gorm.DB, )
