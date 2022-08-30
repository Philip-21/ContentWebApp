package database

import (
	"testing"

	"github.com/Philip-21/Content/models"
	"gorm.io/gorm"
)

func Test_GetContentsdb(t *testing.T, db *gorm.DB) ([]models.Content, error) {
	var cont []models.Content
	return cont, nil
}

func Test_GetContentByIDdb(t *testing.T, id string, db *gorm.DB) (models.Content, bool, error) {
	ID := models.Content{}
	return ID, true, nil
}

func Test_DeleteContentdb(t *testing.T, id string, db *gorm.DB) error {
	return nil
}

func Test_UpdateContentdb(t *testing.T, b *models.Content, db *gorm.DB) error {
	return nil
}

func Test_GetUser(t *testing.T) (models.ContentUser, error) {
	User := models.ContentUser{}
	return User, nil
}
