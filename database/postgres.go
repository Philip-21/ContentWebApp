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

func GetUser(db *gorm.DB) (models.ContentUser, error) {
	user := models.ContentUser{}
	query := db.Select("users.*").Group("users_id")
	err := query.Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserID(id string, db *gorm.DB) (models.ContentUser, error) {
	user := models.ContentUser{}

	query := db.Select("content_users.*")
	query = query.Group("content_users.id")
	err := query.Where("content_users.id=?", id).First(&user).Error
	err1 := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !err1 {
		return user, err
	}
	if err1 {
		return user, nil
	}
	return user, nil
}

func GetUserEmail(email string, db *gorm.DB) (models.ContentUser, error) {
	user := models.ContentUser{}
	choose := db.Select("content_users.*")
	query := choose.Group("content_users.email")
	err := query.Where("content_users.email=?", email).First(&user).Error
	err1 := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !err1 {
		return user, err
	}
	if err1 {
		return user, nil
	}
	return user, nil

}
func GetUserPassword(password string, db *gorm.DB) (models.ContentUser, error) {
	user := models.ContentUser{}
	choose := db.Select("content_users.*")
	query := choose.Group("content_users.hashed_password")
	err := query.Where("content_users.hashed_password=?", password).First(&user).Error
	err1 := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !err1 {
		return user, err
	}
	if err1 {
		return user, nil
	}
	return user, nil

}

//this would be used in implementing login handlers
type AuthUser interface {
	GetUser(db *gorm.DB) (models.ContentUser, error)
	GetUserID(id string, db *gorm.DB) (models.ContentUser, error)
	GetUserEmail(email string, db *gorm.DB) (models.ContentUser, error)
	GetUserPassword(password string, db *gorm.DB) (models.ContentUser, error)
}

//a struct for implementing the interface to be used as a configuration for the user handlers
type Userctx struct {
	AuthUser
}
