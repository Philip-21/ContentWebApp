package database

import (
	"errors"

	"github.com/Philip-21/Content/models"
	"golang.org/x/crypto/bcrypt"
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

func Authenticate(db *gorm.DB, email string, password string) (models.ContentUser, error) {

	user := models.ContentUser{}
	query := db.Select("content_users.*")
	err := query.Where("email= ?", email).Take(&user).Error
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		errors.New("incorrect password")
		//c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("incorrect password %s", req.Password)})
		return user, err
	}

	return user, nil

}
func GetUser(db *gorm.DB, id int, email string, firstname string, lastname string) (models.ContentUser, error) {
	user := models.ContentUser{
		ID:        id,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
	}
	//result := db.Find(&users,"id = ? ")
	err := db.Table("content_users").Select("email", "firstname", "lastname").Where("id = ?", id).Scan(&user)
	if err != nil {
		return user, gorm.ErrRecordNotFound
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
