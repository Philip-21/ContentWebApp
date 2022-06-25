package database

import (
	"errors"

	"github.com/Philip-21/proj1/models"
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

func GetUser(db *gorm.DB) (models.ContentUser, error) {
	user := models.ContentUser{}
	query := db.Select("content_users.*").Group("content_users.id")
	err := query.Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

//[0m←[33m[1158.674ms] ←[34;1m[rows:0]←[0m SELECT content_users.* FROM "content_users" WHERE content_users.id='' AND "content_users"."deleted_at" IS NULL GROUP BY "content_users"."id" ORDER BY "content_users"."id" LIMIT 1

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
func FindByEmail(db *gorm.DB, email string) (models.ContentUser, error) {
	var user models.ContentUser
	res := db.Find(user, &models.ContentUser{Email: email})
	return user, res.Error
}

//this would be used in implementing login handlers
type AuthUser interface {
	GetUserPassword(password string, db *gorm.DB) (models.ContentUser, error)
}

//a struct for implementing the interface to be used as a configuration for the user handlers
type Userctx struct {
	AuthUser
}
