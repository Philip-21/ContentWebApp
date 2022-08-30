package database

import (
	"errors"

	"github.com/Philip-21/Content/models"
	"gorm.io/gorm"
)

func GetContents(db *gorm.DB) ([]models.Content, error) {
	contents := []models.Content{}
	query := db.Select("contents.*").Group("contents.id")
	err := query.Find(&contents).Error
	if err != nil {
		return contents, err
	}
	return contents, nil
}

func GetContentByID(id string, db *gorm.DB) (models.Content, bool, error) {
	c := models.Content{}

	query := db.Select("contents.*")
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
	err := db.Where("id = ? ", id).Delete(&b).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateContent(db *gorm.DB, b *models.Content) error {
	err := db.Save(&b).Error
	if err != nil {
		return err
	}
	return nil
}

//////////Users

func GetUser(db *gorm.DB, email string) (models.ContentUser, error) {

	user := models.ContentUser{}
	query := db.Select("content_users.*")
	err := query.Where("email= ?", email).Take(&user).Error
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return user, err
	}

	return user, nil

}

func UserID(db *gorm.DB, id uint) (models.ContentUser, error) {
	i := models.ContentUser{}
	query := db.Select("content_users.*")
	query = query.Group("content_users.id")
	err := query.Where("content_users.id = ?", id).First(&i).Error
	err1 := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !err1 {
		return i, err
	}
	if err1 {
		return i, nil
	}
	return i, nil
}
